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

    "github.com/wbeebe/rpi/devices"
)

const DEFAULT_ADDRESS int = 0x70

// A very simple test for the Adafruit 8x16 LED Matrix FeatherWing Display.
// "Bounces" a row of lit LEDs from top to bottom and back to the top,
// left to right, leaving a single straight line of lit LEDs across
// the top of the display.
//
func bounce(device i2c.Connection) {
    buffer := make([]byte, 16)
    upDirection := make([]bool, 16)
    altIndex := []int{0,2,4,6,8,10,12,14,1,3,5,7,9,11,13,15}

    for i := range buffer {
        buffer[i] = 0x80
    }

    for i := 0 ; i < 2 * len(buffer) ; i++ {
        device.WriteBlockData(0, buffer)

        if i > 0 {
            time.Sleep(25 * time.Millisecond)
        }

        for j := i ; j >= 0 ; j-- {
            if j < len(buffer) && buffer[altIndex[j]] > 1 && ! upDirection[altIndex[j]] {
                buffer[altIndex[j]] >>= 1
            } else if j < len(buffer) && buffer[altIndex[j]] == 1 {
                upDirection[altIndex[j]] = true
            }

            if j < len(buffer) && buffer[altIndex[j]] < 0x80 && upDirection[altIndex[j]] {
                buffer[altIndex[j]] <<= 1
            }
        }
    }
}

func main() {
    ht16k33 := devices.NewHT16K33Driver(DEFAULT_ADDRESS)

    // Hook the various system abort calls for us to use or ignore as we
    // see fit. In particular hook SIGINT, or CTRL+C for below.
    //
    signal_chan := make(chan os.Signal, 1)
    signal.Notify(signal_chan,
        syscall.SIGHUP,
        syscall.SIGINT,
        syscall.SIGTERM,
        syscall.SIGQUIT)

    err := ht16k33.Start()
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
                ht16k33.Clear()
                ht16k33.Close()
                os.Exit(0)
            default:
            }
        }
    }()

    ht16k33.Clear()

    for i := 0 ; i < 6 ; i++ {
        bounce(ht16k33.Connection())
    }

    time.Sleep(30 * time.Millisecond)

    ht16k33.Clear()
    ht16k33.Close()
}

