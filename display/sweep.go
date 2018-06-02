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

var con i2c.Connection

func initialize() i2c.Connection {
    adapter := raspi.NewAdaptor()
    adapter.Connect()
    bus := adapter.GetDefaultBus()
    fmt.Printf("bus %d\n", bus)
    con, _ = adapter.GetConnection(DEFAULT_ADDRESS, bus)
    // Turn on chip's internal oscillator.
    con.WriteByte(HT16K33_SYSTEM_SETUP | HT16K33_OSCILLATOR_ON)
    // Turn on the display. YOU HAVE TO SEND THIS.
    con.WriteByte(HT16K33_DISPLAY_SETUP | HT16K33_DISPLAY_ON)
    // Set for maximum LED brightness.
    con.WriteByte(HT16K33_CMD_BRIGHTNESS | 0x0f)
    return con
}

// A very simple test for the Adafruit 0.8" 8x16 LED Matrix
// FeatherWing Display.
// Rolls a single bit from top to bottom, left to right, leaving
// a single line of lit bits across the bottom of the display.
//
func lightAll() {
    var bit uint8
    for k := 0 ; k < 2 ; k++ {
        for i := 0 ; i < 8 ; i++ {
            bit = 0x80
            for j := k * 8 ; j < (8 + k*8) ; j++ {
                con.WriteByteData(uint8((i * 2) + k), bit)
                bit >>= 1
                time.Sleep(40 * time.Millisecond)
            }
        }
    }
}

// Just turns every lit LED off.
//
func darkenAll() {
    // Turn off every bit on the displays.
    //
    con.WriteWordData(0, 0)
    con.WriteWordData(2, 0)
    con.WriteWordData(4, 0)
    con.WriteWordData(6, 0)
    con.WriteWordData(8, 0)
    con.WriteWordData(10, 0)
    con.WriteWordData(12, 0)
    con.WriteWordData(14, 0)
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

    con := initialize()

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
                con.Close()
                os.Exit(0)
            default:
            }
        }
    }()

    darkenAll()
    lightAll()
    time.Sleep(1 * time.Second)
    darkenAll()
    con.Close()
}

