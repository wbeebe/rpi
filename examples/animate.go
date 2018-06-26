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
    "sort"
    "time"
    "os"
    "os/signal"
    "syscall"

    "github.com/wbeebe/rpi/devices"
)

const DEFAULT_ADDRESS int = 0x70

// An application for the Adafruit 0.8" 8x16 LED Matrix FeatherWing Display.
//

// Higher level functions.
//
// Simple display glyphs
//
var blockCircle []byte  = []byte {0x3c, 0x42, 0x81, 0x81, 0x81, 0x81, 0x42, 0x3c}
var blockSquare []byte  = []byte {0xFF, 0xFF, 0xC3, 0xC3, 0xC3, 0xC3, 0xFF, 0xFF}
var blockDiamond []byte = []byte {0x18, 0x3C, 0x7E, 0xFF, 0xFF, 0x7E, 0x3C, 0x18}
var blockCheck []byte   = []byte {0xAA, 0x55, 0xAA, 0x55, 0xAA, 0x55, 0xAA, 0x55}
var blockX []byte       = []byte {0x81, 0x42, 0x24, 0x18, 0x18, 0x24, 0x42, 0x81}
var blockFace []byte    = []byte {0x3C, 0x42, 0xA9, 0x89, 0x89, 0xA9, 0x42, 0x3C}
var blockFrown []byte   = []byte {0x3C, 0x42, 0xA5, 0x89, 0x89, 0xA5, 0x42, 0x3C}
var blockSmile []byte   = []byte {0x3C, 0x42, 0xA9, 0x85, 0x85, 0xA9, 0x42, 0x3C}
var blockFslash []byte  = []byte {0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80}
var blockBslash []byte  = []byte {0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01}

// So why define this twice? Because I needed a set to display in insertion order,
// and a map to individually address each glyph by string name.
//
var shapeSet []*[]byte =
   []*[]byte {&blockCircle, &blockSquare, &blockDiamond, &blockCheck, &blockFslash, &blockBslash, &blockX, &blockFace, &blockFrown, &blockSmile}

var shapeTable= map[string]*[]byte{
    "circle": &blockCircle,
    "square": &blockSquare,
    "diamond": &blockDiamond,
    "check": &blockCheck,
    "forwardSlash": &blockFslash,
    "backSlash": &blockBslash,
    "x": &blockX,
    "face": &blockFace,
    "frown": &blockFrown,
    "smile": &blockSmile,
}

func listGlyphNames() {
    var names = make([]string, len(shapeTable))
    index := 0
    for key, _ := range shapeTable {
        names[index] = key
        index++
    }
    sort.Strings(names)
    for _, name := range names {
        fmt.Printf(" %s\n", name)
    }
}

// Simple animation with smiley faces. Similar to what Adafruit shows on their site
// with these 8x16 displays.
//
func simpleAnimation(device *devices.HT16K33Driver) {
    for {
        device.LoadBuffer(*shapeTable["face"], 0)
        device.LoadBuffer(*shapeTable["frown"], 1)
        device.DrawBuffer()
        time.Sleep( 500 * time.Millisecond )
        device.LoadBuffer(*shapeTable["frown"], 0)
        device.LoadBuffer(*shapeTable["smile"], 1)
        device.DrawBuffer()
        time.Sleep( 500 * time.Millisecond )
        device.LoadBuffer(*shapeTable["smile"], 0)
        device.LoadBuffer(*shapeTable["face"], 1)
        device.DrawBuffer()
        time.Sleep( time.Second )
    }
}

// Scroll's two glyphs across the display.
//
func simpleScroll(device *devices.HT16K33Driver, glyphName string) {
    device.LoadBuffer(*shapeTable[glyphName], 0)
    device.LoadBuffer(*shapeTable[glyphName], 1)

    for {
        device.DrawBuffer()
        time.Sleep( 250 * time.Millisecond )
        device.RotateBuffer()
    }
}

// Displays a simple triangle wave across the display.
//
func wave(device *devices.HT16K33Driver, cycles int) {
    device.LoadBuffer(blockBslash, 0)
    device.LoadBuffer(blockFslash, 1)

    for c := 0 ; c < cycles ; c++ {
        for i := 0 ; i < 16 ; i++ {
            device.DrawBuffer()
            time.Sleep( 30 * time.Millisecond)
            device.RotateBuffer()
        }
    }
}

func shapes(device *devices.HT16K33Driver) {
    for _, glyph := range shapeSet {
        device.LoadBuffer(*glyph, 0)
        device.LoadBuffer(*glyph, 1)
        device.DrawBuffer()
        time.Sleep( 500 * time.Millisecond)
    }
}

func vt52(device *devices.HT16K33Driver) {
    for _, char := range devices.VT52rom {
        device.LoadBuffer(char, 0)
        device.LoadBuffer(char, 1)
        device.DrawBuffer()
        time.Sleep( 350 * time.Millisecond)
    }
}

func help() {
    helpText := []string {
        "\n Adafruit 8x16 Featherwing Display utility\n",
        " Command line actions:\n",
        " faces  - Displays a series of three smiley faces.",
        " shapes - Displays a series of simple glyphs.",
        " scroll - Scrolls a selected glyph from left to right.",
        "        - scroll by itself scrolls a smiley face.",
        "        - 'animate scroll list' lists all glyphs.",
        " vt52   - Displays all the old VT-52 ROM characters",
        "        - translated to work with the Adafruit display.",
        " wave   - Displays a scrolling triangle wave for 10 cycles.\n",
        " No command - this help\n",
    }

    for _, line := range helpText {
        fmt.Println(line)
    }
}

//
//
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

    var action, argument string

    if len(os.Args) > 1 {
        action = os.Args[1]
    }
    if len(os.Args) > 2 {
        argument = os.Args[2]
    }

    switch action {
    case "faces":
        simpleAnimation(ht16k33)
    case "scroll":
        if len(argument) == 0 {
            argument = "smile"
        } else {
            if argument == "list" {
                fmt.Printf("\n Scrollable glyphs are named:\n\n")
                listGlyphNames()
                break
            }
            if _, exist := shapeTable[argument]; ! exist {
                fmt.Printf("\n Glyph %s does not exist.\n", argument)
                fmt.Printf(" Please use one of:\n\n")
                listGlyphNames()
                break
            }
        }

        simpleScroll(ht16k33, argument)
    case "wave":
        wave(ht16k33, 10)
    case "shapes":
        shapes(ht16k33)
    case "vt52":
        vt52(ht16k33)
    default:
        help()
    }

    ht16k33.Clear()
    ht16k33.Close()
}

