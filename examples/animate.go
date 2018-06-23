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

    "gobot.io/x/gobot/drivers/i2c"
    "gobot.io/x/gobot/platforms/raspi"

    "github.com/wbeebe/rpi/devices"
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

// An application for the Adafruit 0.8" 8x16 LED Matrix FeatherWing Display.
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
var blockCheck []byte   = []byte {0xAA, 0x55, 0xAA, 0x55, 0xAA, 0x55, 0xAA, 0x55}
var blockX []byte       = []byte {0x81, 0x42, 0x24, 0x18, 0x18, 0x24, 0x42, 0x81}
var blockFace []byte    = []byte {0x3C, 0x42, 0xA9, 0x89, 0x89, 0xA9, 0x42, 0x3C}
var blockFrown []byte   = []byte {0x3C, 0x42, 0xA5, 0x89, 0x89, 0xA5, 0x42, 0x3C}
var blockSmile []byte   = []byte {0x3C, 0x42, 0xA9, 0x85, 0x85, 0xA9, 0x42, 0x3C}

// So why define this twice? Because I needed a set to display in insertion order,
// and a map to individually address each glyph by string name.
//
var shapeSet []*[]byte =
   []*[]byte {&blockCircle, &blockSquare, &blockDiamond, &blockCheck, &blockX, &blockFace, &blockFrown, &blockSmile}

var shapeTable= map[string]*[]byte{
    "circle": &blockCircle,
    "square": &blockSquare,
    "diamond": &blockDiamond,
    "check": &blockCheck,
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
func simpleAnimation(device i2c.Connection) {
    for {
        loadBuffer(*shapeTable["face"], 0)
        loadBuffer(*shapeTable["frown"], 1)
        drawBuffer(device)
        time.Sleep( 500 * time.Millisecond )
        loadBuffer(*shapeTable["frown"], 0)
        loadBuffer(*shapeTable["smile"], 1)
        drawBuffer(device)
        time.Sleep( 500 * time.Millisecond )
        loadBuffer(*shapeTable["smile"], 0)
        loadBuffer(*shapeTable["face"], 1)
        drawBuffer(device)
        time.Sleep( time.Second )
    }
}

// Scroll's two glyphs across the display.
//
func simpleScroll(device i2c.Connection, glyphName string) {
    loadBuffer(*shapeTable[glyphName], 0)
    loadBuffer(*shapeTable[glyphName], 1)

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
    for _, glyph := range shapeSet {
        loadBuffer(*glyph, 0)
        loadBuffer(*glyph, 1)
        drawBuffer(device)
        time.Sleep( 500 * time.Millisecond)
    }
}

func vt52(device i2c.Connection) {
    for _, char := range devices.VT52rom {
        loadBuffer(char, 0)
        loadBuffer(char, 1)
        drawBuffer(device)
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

    var action, argument string

    if len(os.Args) > 1 {
        action = os.Args[1]
    }
    if len(os.Args) > 2 {
        argument = os.Args[2]
    }

    switch action {
    case "faces":
        simpleAnimation(device)
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

        simpleScroll(device, argument)
    case "wave":
        wave(device)
    case "shapes":
        shapes(device)
    case "vt52":
        vt52(device)
    default:
        help()
    }

    darkenAll(device)
    device.Close()
}

