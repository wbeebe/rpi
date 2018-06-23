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

    "github.com/wbeebe/rpi/devices"
)

func help() {
    helpText := []string {
        "\nHT16K33 Test Utility",
        "For the Adafruit Quad Alphanumeric FeatherWing Display\n",
        "Command line actions:",
        "  bit #    - Takes a bit pattern in binary format, up to 16 bits long, and displays it on a single digit.",
        "           - Leading binary zeros are not necessary.",
        "           - 0000001010111011, which displays '@', and 1010111011 are equivalent.",
        "  clear    - Clears all four characters and turns off all segments.",
        "           - Useful for turning off randomly lit segments while experimenting.",
        "  numbers  - Counts from 0 to F simultaniously in all four digits.",
        "  scroll   - Scrolls a message string passed as a second argument.",
        "           - Messages with spaces will need to be quoted.",
        "  segments - Lights all segments individually including decimal point.",
        "           - Outer segments are lit, then inner.",
        "           - Hex value is displayed in first two digits, third digit displays corresponding individually lit segment.",
        "  table    - Scrolls all defined alphanumeric entries in the internal mapping table across the display, right to left.",
        "  test     - Fully tests all four characters, one at a time, left to right.",
        "           - All segments, including decimal point, are lit.",
        "No command - this help\n",
        " Examples:",
        " display bit 0000001010111011",
        " display scroll \"The quick brown fox\"",
        " display test",
        " display clear\n",
    }

    for _, line := range helpText {
        fmt.Println(line)
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
                devices.Close()
                os.Exit(0)
            default:
            }
        }
    }()

    // Execute actions passed on the command line, along with any additional arguments.
    //
    var action, argument string

    if len(os.Args) > 1 {
        action = os.Args[1]
    }
    if len(os.Args) == 3 {
        argument = os.Args[2]
    }

    switch action {
    case "bit":
        if len(argument) == 0 {
            fmt.Printf(" bit command needs a binary argument.\n")
        } else {
            devices.DisplayBinary(argument)
        }
    case "clear":
        devices.Clear()
    case "numbers":
        devices.NumbersTest()
    case "segments":
        devices.CycleSegments()
    case "scroll":
        if len(argument) == 0 {
            fmt.Printf(" scroll command needs a message to display.\n")
        } else {
            devices.ScrollString(argument)
        }
    case "table":
        devices.ScrollAlphaTable()
    case "test":
        devices.AllDigitSegmentTest()
    default:
        help()
    }
}

