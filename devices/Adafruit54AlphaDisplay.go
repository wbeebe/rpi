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

package devices

import (
    "fmt"
    "sort"
    "strconv"
    "strings"
    "time"
)

// A map of ASCII characters in string format to bit maps to display
// that character.
//
// Mappings are originally from
// https://github.com/adafruit/Adafruit_LED_Backpack/blob/master/Adafruit_LEDBackpack.cpp
// and suitably modified for Go.
//
var alphaTable = map[string]uint16 {
    " ":0x0000,
    "!":0x0006,
    "\"":0x0220,
    "#":0x12CE,
    "$":0x12ED,
    "%":0x0C24,
    "&":0x235D,
    "'":0x0400,
    "(":0x2400,
    ")":0x0900,
    "*":0x3FC0,
    "+":0x12C0,
    ",":0x0800,
    "-":0x00C0,
    ".":0x3808,
    "/":0x0C00,
    "0":0x0C3F,
    "1":0x0006,
    "2":0x00DB,
    "3":0x008F,
    "4":0x00E6,
    "5":0x2069,
    "6":0x00FD,
    "7":0x0007,
    "8":0x00FF,
    "9":0x00EF,
    ":":0x0009,
    ";":0x0A00,
    "<":0x2400,
    "=":0x00C8,
    ">":0x0900,
    "?":0x1083,
    "@":0x02BB,
    "A":0x00F7,
    "B":0x128F,
    "C":0x0039,
    "D":0x120F,
    "E":0x00F9,
    "F":0x0071,
    "G":0x00BD,
    "H":0x00F6,
    "I":0x1200,
    "J":0x001E,
    "K":0x2470,
    "L":0x0038,
    "M":0x0536,
    "N":0x2136,
    "O":0x003F,
    "P":0x00F3,
    "Q":0x203F,
    "R":0x20F3,
    "S":0x00ED,
    "T":0x1201,
    "U":0x003E,
    "V":0x0C30,
    "W":0x2836,
    "X":0x2D00,
    "Y":0x1500,
    "Z":0x0C09,
    "[":0x0039,
    "\\":0x2100,
    "]":0x000F,
    "^":0x0C03,
    "_":0x0008,
    "`":0x0100,
    "a":0x1058,
    "b":0x2078,
    "c":0x00D8,
    "d":0x088E,
    "e":0x0858,
    "f":0x0071,
    "g":0x048E,
    "h":0x1070,
    "i":0x1000,
    "j":0x000E,
    "k":0x3600,
    "l":0x0030,
    "m":0x10D4,
    "n":0x1050,
    "o":0x00DC,
    "p":0x0170,
    "q":0x0486,
    "r":0x0050,
    "s":0x2088,
    "t":0x0078,
    "u":0x001C,
    "v":0x2004,
    "w":0x2814,
    "x":0x28C0,
    "y":0x200C,
    "z":0x0848,
    "{":0x0949,
    "|":0x1200,
    "}":0x2489,
    "~":0x0520,
    "°":0x00E3,
}

type Adafruit54AlphaDisplay struct {
    name string
    ht16k33 *HT16K33Driver
    neighborDisplay *Adafruit54AlphaDisplay
    value1, value2, value3, value4 uint16
}

func NewAdafruit54AlphaDisplay(ht *HT16K33Driver) *Adafruit54AlphaDisplay {
    alpha := &Adafruit54AlphaDisplay {
        name: "Adafruit54AlphaDisplay",
        ht16k33: ht,
    }

    return alpha
}

func (d *Adafruit54AlphaDisplay) Name() string { return d.name }
func (d *Adafruit54AlphaDisplay) SetName(newName string ) { d.name = newName }
func (d *Adafruit54AlphaDisplay) HT16K33() *HT16K33Driver { return d.ht16k33 }
func (d *Adafruit54AlphaDisplay) NeighborDisplay() *Adafruit54AlphaDisplay { return d.neighborDisplay }
func (d *Adafruit54AlphaDisplay) SetNeighborDisplay(nd *Adafruit54AlphaDisplay) { d.neighborDisplay = nd }

func (d *Adafruit54AlphaDisplay) CountDeviceDigits() int {
    var count int = 4

    if d.neighborDisplay != nil {
        count += d.neighborDisplay.CountDeviceDigits()
    }

    return count
}

// Write a 16-bit value to one of the digits.
// Maximum value is 0x7FFF, which turns on all segments and the
// decimal point.
//
func (d *Adafruit54AlphaDisplay) RawWriteDigit(digit uint8, val uint16) {
    display := d.ht16k33.Connection()
    if display != nil {
        display.WriteWordData(digit * 2, val)
    }
}

// A basic function to clear all the digit's backing memory and turn off
// all segments and the decimal point on all digits.
//
func (d *Adafruit54AlphaDisplay) Clear() {
    if d.neighborDisplay != nil {
        d.neighborDisplay.Clear()
    }

    d.RawWriteDigit(0,0)
    d.RawWriteDigit(1,0)
    d.RawWriteDigit(2,0)
    d.RawWriteDigit(3,0)
}

// A basic function to display any hex number from 0 to F.
//
func (d *Adafruit54AlphaDisplay) DisplayNumber(digit uint8, val uint8) {
    if val < 16 {
        d.RawWriteDigit(digit, alphaTable[fmt.Sprintf("%X",val)])
    }
}

// Will display the value of a byte on two consecutive digits.
//
func (d *Adafruit54AlphaDisplay) DisplayByte(digit uint8, val uint8) {
        lowNibble := val & 0xF
        highNibble := val  >> 4
        d.DisplayNumber( digit, highNibble)
        d.DisplayNumber( digit + 1, lowNibble)
}

// A test to cycle through lighting all the segments plus decimal point on a given digit.
//
func (d *Adafruit54AlphaDisplay) CycleDigit(digit uint8) {
    d.RawWriteDigit(digit, 0x7fff)
    time.Sleep(500 * time.Millisecond)
    d.RawWriteDigit(digit, 0x00ff)
    time.Sleep(500 * time.Millisecond)
    d.RawWriteDigit(digit, 0x7f00)
    time.Sleep(500 * time.Millisecond)
    d.RawWriteDigit(digit, 0)
}

// A test function to drive CycleDigit for all digits.
//
func (d *Adafruit54AlphaDisplay) AllDigitSegmentTest() {
    d.Clear()
    if d.neighborDisplay != nil { d.neighborDisplay.AllDigitSegmentTest() }
    d.CycleDigit(0)
    d.CycleDigit(1)
    d.CycleDigit(2)
    d.CycleDigit(3)
}

// A test to cycle through each segment in a digit.
// The hex value that enables each segment is displayed
// alongside the lit segment on another digit.
// Because there are only four segments, the bit value that
// is sent to rawWriteDigit() is split into a low byte,
// then a high byte to display it's hexadecimal value in
// just two digits to the left.
//
func (d *Adafruit54AlphaDisplay) CycleSegments() {
    d.Clear()
    var bit uint16
    var i int
    var digit uint8
    bit = 1

    for i = 0 ; i < 16 ; i++ {
        d.RawWriteDigit( 2, bit)
        if i < 8 {
            digit = uint8(bit)
        } else {
            digit = uint8(bit >> 8)
        }
        d.DisplayByte(0, digit)
        bit <<= 1
        time.Sleep(500 * time.Millisecond)
    }

    d.Clear()
}

// A test function to display hexademical numbers simultaniously
// on all digits.
//
func (d *Adafruit54AlphaDisplay) NumbersTest() {
    d.Clear()
    if d.neighborDisplay != nil { d.neighborDisplay.NumbersTest() }
    var i uint8
    for i = 0 ; i < 16 ; i++ {
        d.DisplayNumber(0, i)
        d.DisplayNumber(1, i)
        d.DisplayNumber(2, i)
        d.DisplayNumber(3, i)
        time.Sleep(500 * time.Millisecond)
    }
    d.Clear()
}

// Will take a binary representation in the form 0000000000000000
// and convert it into a unsigned binary int to display on a single
// digit. Part of the bit command, and a good way to see how to light
// any combination of segments for testing and simple investigation.
//
func (d *Adafruit54AlphaDisplay) DisplayBinary(digit string) {
    val, err := strconv.ParseUint(digit, 2, 16)

    if err != nil {
        fmt.Printf(" Invalid bit argument: %s\n", digit)
        return
    }

    d.RawWriteDigit(2, uint16(val))
}

// A test function to scroll the contents of the alpha table across the display.
// The scroll is in ascending sorted order.
//
func (d *Adafruit54AlphaDisplay) ScrollAlphaTable() {
    var keys []string

    for key := range alphaTable {
        keys = append(keys, key)
    }

    sort.Strings(keys)
    d.ScrollString(strings.Join(keys, ""))
}

// Scroll an alphnumeric string across the digits.
//
func (d *Adafruit54AlphaDisplay) ScrollString(message string) {
    var valueOut, value1, value2, value3, value4 uint16

    for _, key := range message {
        valueOut = value1
        value1 = value2
        value2 = value3
        value3 = value4
        value4 = alphaTable[string(key)]
        d.RawWriteDigit(0, value1)
        d.RawWriteDigit(1, value2)
        d.RawWriteDigit(2, value3)
        d.RawWriteDigit(3, value4)
        if d.neighborDisplay != nil {
            d.neighborDisplay.ScrollInFromRight(valueOut)
        }
        time.Sleep(400 * time.Millisecond)
    }

    time.Sleep(1 * time.Second)
    d.Clear()
}

func (d *Adafruit54AlphaDisplay) ScrollInFromRight(incoming uint16) {
    d.value1 = d.value2
    d.value2 = d.value3
    d.value3 = d.value4
    d.value4 = incoming
    d.RawWriteDigit(0, d.value1)
    d.RawWriteDigit(1, d.value2)
    d.RawWriteDigit(2, d.value3)
    d.RawWriteDigit(3, d.value4)
}

func (d *Adafruit54AlphaDisplay) WriteDirect(message string) {
    digits := d.CountDeviceDigits()

    if len(message) > digits { message = message[0:digits] }

    if d.neighborDisplay != nil {
        lim := len(message) - 4
        if lim > 0 { d.neighborDisplay.WriteDirect(message[0:lim]) }
    }

    var cindex uint8 = uint8(4 - len(message))
    for _, letter := range message {
        d.RawWriteDigit(cindex, alphaTable[string(letter)])
        cindex += 1
    }
}

// Essentially a wrapper for i2c.Connection.Close()
// with a call to clear the display first.
// Call this last before exiting an application.
//
func (d *Adafruit54AlphaDisplay) Close() {
    if d.neighborDisplay != nil {
        d.neighborDisplay.Clear()
        d.neighborDisplay.HT16K33().Close()
    }
    d.Clear()
    d.ht16k33.Close()
}
