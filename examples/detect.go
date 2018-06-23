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
    "os"
    "os/signal"
    "syscall"

    "gobot.io/x/gobot/platforms/raspi"
)


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
                os.Exit(0)
            default:
            }
        }
    }()

    // Go find an I2C buss and open it.
    //
    adapter := raspi.NewAdaptor()
    adapter.Connect()
    bus := adapter.GetDefaultBus()

    // Now iterate across all I2C device addresses.
    // If we successfully read a byte from an address,
    // then we have detected a device. Print out the
    // hex address of that device.
    //
    for i := 0 ; i < 128 ; i++ {
        if device, err := adapter.GetConnection(i, bus) ; err == nil {
            if _, err := device.ReadByte() ; err == nil {
                fmt.Printf( " Found device at 0x%x / %d on I2C bus %d\n", i, i, bus)
            }
        }
    }
}

