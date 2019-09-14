/*
Copyright (c) 2018 William H. Beebe, Jr.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// An example application for driving either a DL2416 or DL1414
// intelligent display with the MCP23017 I2C port expander on
// a Raspberry Pi 3 B Plus. Capable of driving up to four
// individual intelligent displays as coded and wired.
//
package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "gobot.io/x/gobot/drivers/i2c"
    "gobot.io/x/gobot/platforms/raspi"
)

const MCP23017_DEFAULT_ADDRESS int = 0x20
var address int = MCP23017_DEFAULT_ADDRESS

// Go's iota is used to define MCP23017 register addresses.
// The underscore is used to skip undefined registers addresses.
// These are register addresses in bank 0.
//
const (
    IODIRA = iota
    IODIRB
    IPOLA
    IPOLB
    GPINTENA
    GPINTENB
    _
    _
    _
    _
    _
    _
    GPPUA
    GPPUB
    _
    _
    _
    _
    GPIOA
    GPIOB
    OLATA
    OLATB
)

func initialize() (device i2c.Connection, err error) {
    adapter := raspi.NewAdaptor()
    adapter.Connect()
    bus := adapter.GetDefaultBus()

    // Check to see if the device actually is on the I2C buss.
    // If it is then use it, else return an error.
    //
    if device, err := adapter.GetConnection(address, bus) ; err == nil {
        if _, err := device.ReadByte() ; err == nil {
            fmt.Printf(" Using device 0x%x / %d on bus %d\n", address, address, bus)
        } else {
            return device, fmt.Errorf(" Could not find device 0x%x / %d", address, address)
        }
    }

    device, _ = adapter.GetConnection(address, bus)
    // Set the GPIO ports (A and B) to all outputs
    //
    device.WriteByteData(IODIRA, 0x00)
    device.WriteByteData(IODIRB, 0x00)
    return device, nil
}

// writeDisplayChar allows for up to four blocks of intelligent displays,
// with four characters/block, for a total of 16 characters.
// The individual characters are addressed as 0-15,
// starting from the right and going left:
//
// (15) (14) (13) (12) (11) (10) (9) (8) (7) (6) (5) (4) (3) (2) (1) (0)
// | block 3          | block 2         | block 1       | block 0       |
//
// Very simple math and a switch statement build up the necessary combined
// block write/enable and character address within the block, which is then
// written out as a single byte to GPIO A, the addressing GPIO.
//
// GPIO A is split logically into two nibbles. The lower nibble emits the
// individual character address from 0 to 3. The upper nibble emits the
// combined device /WR/CE signal, where low writes/selects the device and
// high deselects the device.
//
// The upper nibble device write/select bit pattern is:
//
// GPIOA.7 - block 3
// GPIOA.6 - block 2
// GPIOA.5 - block 1
// GPIOA.4 - block 0
//
// Two GPIO A output lines are current unused in this scheme, 2 and 3.
//
const DISPLAY_MAX_CHARS int = 4

func writeDisplayChar(device i2c.Connection, location int, char uint8) {
    digit := location % DISPLAY_MAX_CHARS
    block := location / DISPLAY_MAX_CHARS

    // With the character/block address calculated, write out the char
    // data. Then OR the address with 0xF0 to make the select bits all
    // high, writing the data into the individual character.
    // This works because on all devices the /WR line is in essence
    // the /CE line as well.
    //
    digitAddress := [4]uint8{0xE0, 0xD0, 0xB0, 0x70}
    device.WriteByteData(GPIOA, 0xF0)
    device.WriteByteData(GPIOA, (digitAddress[block] | uint8(digit)))
    device.WriteByteData(GPIOB, char)
    device.WriteByteData(GPIOA, 0xF0)
}

// A very basic, raw write function that addresses all display
// blocks at once, writing a space to each digit, turning off
// any lit segments.
//
func clearDisplay(device i2c.Connection) {
    device.WriteByteData(GPIOA, 0x03)
    device.WriteByteData(GPIOB, ' ')
    device.WriteByteData(GPIOA, 0x02)
    device.WriteByteData(GPIOB, ' ')
    device.WriteByteData(GPIOA, 0x01)
    device.WriteByteData(GPIOB, ' ')
    device.WriteByteData(GPIOA, 0x00)
    device.WriteByteData(GPIOB, ' ')
}

// A very basic text scrolling routine that assumes
// four blocks/16 characters are to be written to.
// If the text string length is less than the maximum
// characters then the string is just displayed across
// all the devices.
//
func scrollText(device i2c.Connection, text string) {
    clearDisplay(device)
    MAX_CHARS := DISPLAY_MAX_CHARS * 4

    if len(text) < MAX_CHARS {
        offset := 0
        for j := 15 ; j >= MAX_CHARS - len(text) ; j-- {
            writeDisplayChar(device, j, text[offset])
            offset++
        }

        return
    }

    for i := 0 ; i <= len(text) - MAX_CHARS; i++ {
        offset := 0
        for j := 15 ; j >= 0 ; j-- {
            writeDisplayChar(device, j, text[i + offset])
            offset++
        }

        time.Sleep( 250 * time.Millisecond)
    }
}

// A very basic clock. Writes the time to the
// two lowest blocks (characters 0-7) and the date to the
// two highest blocks (character 8-15)
//
func basicClock(device i2c.Connection) {
    for {
        now := time.Now()
        text := now.String()
        ti := 11 // time index into string
        di := 2  // date index into string

        for i := 7 ; i >= 0 ; i-- {
            writeDisplayChar(device, i, text[ti])
            writeDisplayChar(device, i+8, text[di])
            ti++
            di++
        }

        time.Sleep( 1 * time.Second)
    }
}

func main() {
    // Hook the various system abort calls for us to use or ignore as we
    // see fit. In particular hook SIGINT, or CTRL+C for below.
    //
    signal_chan := make(chan os.Signal, 1)
    signal.Notify(signal_chan,
        syscall.SIGHUP,
        syscall.SIGINT,
        syscall.SIGTERM,
        syscall.SIGQUIT)

    device, err := initialize()
    if err != nil {
        log.Fatal(err)
    }

    // We want to capture CTRL+C to first clear the display and then exit.
    // We don't want to leave the display lit on an abort.
    //
    go func() {
        for {
            signal := <-signal_chan
            switch signal {
            case syscall.SIGINT:
                // CTRL+C
                fmt.Println()
                clearDisplay(device)
                device.Close()
                os.Exit(0)
            default:
            }
        }
    }()

    scrollText(device, "CHARACTER TEST")
    time.Sleep( 3 * time.Second)

    testchars := "                !\"#$%&'<>*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_                "
    scrollText(device, testchars)

    basicClock(device)
}

