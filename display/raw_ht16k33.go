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
    "strconv"
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

var address = DEFAULT_ADDRESS
var device i2c.Connection

func initialize() (device i2c.Connection, err error) {
    adapter := raspi.NewAdaptor()
    adapter.Connect()
    bus := adapter.GetDefaultBus()

    // Check to see if the device actually is on the I2C buss.
    // If it is then use it, else return an error.
    //
    if device, err := adapter.GetConnection(address, bus) ; err == nil {
        if _, err := device.ReadByte() ; err == nil {
            fmt.Printf(" Using device 0x%x/%d on bus %d\n", address, address, bus)
        } else {
            return device, fmt.Errorf(" Could not find device 0x%x / %d", address, address)
        }
    }

    device, _ = adapter.GetConnection(address, bus)
    // Turn on chip's internal oscillator.
    device.WriteByte(HT16K33_SYSTEM_SETUP | HT16K33_OSCILLATOR_ON)
    // Turn on the display. YOU HAVE TO SEND THIS.
    device.WriteByte(HT16K33_DISPLAY_SETUP | HT16K33_DISPLAY_ON)
    // Set for maximum LED brightness.
    device.WriteByte(HT16K33_CMD_BRIGHTNESS | 0x0f)
    return device, nil
}

func lightAll() {
    // First four digits for Alphanumeric and 8x16 Matrix
    // FeatherWing Displays.
    //
    // The 'digit' address is the address/offset into the
    // HT16K33's internal eight byte array. Each bit
    // represents a segment or LED, each address a section
    // within an entire device
    //
    // Digit 0
    //
    device.WriteWordData(0, 0xFFFF)

    // Digit 1
    //
    device.WriteWordData(2, 0xFFFF)

    // Digit 2
    //
    device.WriteWordData(4, 0xFFFF)

    // Digit 3
    //
    device.WriteWordData(6, 0xFFFF)

    // Rest of the bytes for the
    // Adafruit 0.8" 8x16 LED Matrix FeatherWing Display
    device.WriteWordData(8, 0xFFFF)
    device.WriteWordData(10, 0xFFFF)
    device.WriteWordData(12, 0xFFFF)
    device.WriteWordData(14, 0xFFFF)
}

func darkenAll() {
    // Turn off every segment on every digit.
    //
    var data []byte = make([]byte, 16)
    device.WriteBlockData(0, data)
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

    for _, arg := range os.Args[1:] {
        if newAddress, err := strconv.ParseInt(arg, 0, 32); err == nil {
            address = int(newAddress)
        } else {
            fmt.Println(err)
        }
    }

    dev, err := initialize()
    if err != nil {
        log.Fatal(err)
    }
    device = dev

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
                darkenAll()
                device.Close()
                os.Exit(0)
            default:
            }
        }
    }()

    darkenAll()
    lightAll()
    time.Sleep(5 * time.Second)
    darkenAll()
    device.Close()
}

