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

package main

import (
    "fmt"
    "log"
    "time"
    "os"
    "os/signal"
    "syscall"

    "gobot.io/x/gobot/drivers/i2c"
    "gobot.io/x/gobot/platforms/raspi"
)

const DEFAULT_ADDRESS int = 0x70

// Commands to send the HT16K33
//
const (
    HT16K33_SYSTEM_SETUP byte = 0x20
    HT16K33_OSCILLATOR_ON byte = 0x01
    HT16K33_DISPLAY_SETUP byte = 0x80
    HT16K33_DISPLAY_ON byte = 0x01
    HT16K33_BLINK_OFF byte = 0x00
    HT16K33_BLINK_2HZ byte = 0x02
    HT16K33_BLINK_1HZ byte = 0x04
    HT16K33_BLINK_HALFHZ byte = 0x06
    HT16K33_CMD_BRIGHTNESS byte = 0xE0
)

func initialize(address int) (device i2c.Connection, err error) {
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

    device, _ = adapter.GetConnection(DEFAULT_ADDRESS, bus)
    // Turn on chip's internal oscillator.
    device.WriteByte(HT16K33_SYSTEM_SETUP | HT16K33_OSCILLATOR_ON)
    // Turn on the display. YOU HAVE TO SEND THIS.
    device.WriteByte(HT16K33_DISPLAY_SETUP | HT16K33_DISPLAY_ON)
    // Set for maximum LED brightness.
    device.WriteByte(HT16K33_CMD_BRIGHTNESS | 0x0f)
    return device, nil
}

// A application for the Adafruit 0.8" 8x16 LED Matrix FeatherWing Display.
//

var buffer []byte = make([]byte, 16)
var altIndex []int = []int{0,2,4,6,8,10,12,14,1,3,5,7,9,11,13,15}

// Loads the buffer with data, in the pattern necessary for proper
// displaying. Works with the concept of blocks that matches the
// 8x8 LED arrays on the display. Block 0 is on the left, block 1
// on the right.
//
func loadBuffer(bits []byte, block int) {
    block &= 0x01

    for i := 0; i < len(bits) ; i++ {
        buffer[altIndex[i + block * 8]] = bits[i]
    }
}

// A wrapper for WriteBlockData for displaying the buffer.
//
func drawBuffer(device i2c.Connection) {
    device.WriteBlockData(0, buffer)
}

// Rotates the buffer contents from left to right.
//
func rotateBuffer() {
    end := buffer[altIndex[len(buffer) - 1]]

    for i := len(buffer) - 1 ; i > 0 ; i-- {
        buffer[altIndex[i]] = buffer[altIndex[i-1]]
    }

    buffer[0] = end
}

// Turns every lit LED off by writing binary zeros to all locations.
//
func darkenAll(device i2c.Connection) {
    block := make([]byte, 16)
    device.WriteBlockData(0, block)
}

// Higher level functions.
//
// Simple display glyphs
//
var blockCircle []byte  = []byte {0x3c, 0x42, 0x81, 0x81, 0x81, 0x81, 0x42, 0x3c}
var blockSquare []byte  = []byte {0xFF, 0xFF, 0xC3, 0xC3, 0xC3, 0xC3, 0xFF, 0xFF}
var blockDiamond []byte = []byte {0x18, 0x3C, 0x7E, 0xFF, 0xFF, 0x7E, 0x3C, 0x18}
var blockX []byte       = []byte {0x81, 0x42, 0x24, 0x18, 0x18, 0x24, 0x42, 0x81}
var blockFace []byte    = []byte {0x3C, 0x42, 0xA9, 0x89, 0x89, 0xA9, 0x42, 0x3C}
var blockFrown []byte   = []byte {0x3C, 0x42, 0xA5, 0x89, 0x89, 0xA5, 0x42, 0x3C}
var blockSmile []byte   = []byte {0x3C, 0x42, 0xA9, 0x85, 0x85, 0xA9, 0x42, 0x3C}

var shapeTable []*[]byte =
    []*[]byte {&blockCircle, &blockSquare, &blockDiamond, &blockX, &blockFace, &blockFrown, &blockSmile}

// Simple animation with smiley faces. Similar to what Adafruit shows on their site
// with these 8x16 displays.
//
func simpleAnimation(device i2c.Connection) {
    for {
        loadBuffer(blockFace, 0)
        loadBuffer(blockFrown, 1)
        drawBuffer(device)
        time.Sleep( 500 * time.Millisecond )
        loadBuffer(blockFrown, 0)
        loadBuffer(blockSmile, 1)
        drawBuffer(device)
        time.Sleep( 500 * time.Millisecond )
        loadBuffer(blockSmile, 0)
        loadBuffer(blockFace, 1)
        drawBuffer(device)
        time.Sleep( time.Second )
    }
}

// Scroll's two smiley face glyphs across the display.
//
func simpleScroll(device i2c.Connection) {
    loadBuffer(blockSmile, 0)
    loadBuffer(blockFrown, 1)

    for {
        drawBuffer(device)
        time.Sleep( 250 * time.Millisecond )
        rotateBuffer()
    }
}

// Displays a simple triangle wave across the display.
//
func wave(device i2c.Connection) {
    var bit byte = 1
    var blen int= len(buffer)

    for i := 0 ; i < blen/2 ; i++ {
        buffer[altIndex[i]] = bit
        buffer[altIndex[blen - 1 - i]] = bit
        bit <<= 1
    }

    for {
        device.WriteBlockData(0, buffer)
        time.Sleep( 30 * time.Millisecond)
        rotateBuffer()
    }
}

func shapes(device i2c.Connection) {
    for _, glyph := range shapeTable {
        loadBuffer(*glyph, 0)
        loadBuffer(*glyph, 1)
        device.WriteBlockData(0, buffer)
        time.Sleep( 500 * time.Millisecond)
    }
}

func help() {
    helpText := []string {
        "\n Adafruit 8x16 Featherwing Display utility\n",
        " Command line actions:\n",
        " faces  - Displays a series of three smiley faces.",
        " shapes - Displays a series of simple glyphs.",
        " scroll - Scrolls smiley faces from left to right.",
        " wave   - Displays a scrolling triangle wave.\n",
        " No command - this help\n",
    }

    for _, line := range helpText {
        fmt.Println(line)
    }
}

//
//
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

    device, err := initialize(DEFAULT_ADDRESS)
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
                darkenAll(device)
                device.Close()
                os.Exit(0)
            default:
            }
        }
    }()

    var action string

    if len(os.Args) > 1 {
        action = os.Args[1]
    }

    switch action {
    case "faces":
        simpleAnimation(device)
    case "scroll":
        simpleScroll(device)
    case "wave":
        wave(device)
    case "shapes":
        shapes(device)
    default:
        help()
    }

    darkenAll(device)
    device.Close()
}

