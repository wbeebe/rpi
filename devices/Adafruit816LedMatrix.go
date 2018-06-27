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

type Adafruit816LedMatrix struct {
    name string
    ht16k33 *HT16K33Driver
    buffer []byte
    altIndex []int
}

func NewAdafruit816LedMatrix(ht *HT16K33Driver) *Adafruit816LedMatrix {
    matrix := &Adafruit816LedMatrix {
        name: "Adafruit816LedMatrixDisplay",
        ht16k33: ht,
    }

    // Matches the HT16K33's internal data buffer.
    //
    matrix.buffer = make([]byte, 16)

    // A simple way to map a regular index to the
    // interleaved layout that is wired to various
    // LEDs on various Adafruit displays.
    //
    matrix.altIndex = []int{0,2,4,6,8,10,12,14,1,3,5,7,9,11,13,15}

    return matrix
}

func (d *Adafruit816LedMatrix) Name() string { return d.name }
func (d *Adafruit816LedMatrix) SetName(newName string ) { d.name = newName }
func (d *Adafruit816LedMatrix) HT16K33() *HT16K33Driver { return d.ht16k33 }

// Loads the buffer with data, in the pattern necessary for proper
// displaying. Works with the concept of blocks that matches the
// 8x8 LED arrays on the display. Block 0 is on the left, block 1
// on the right.
//
func (d *Adafruit816LedMatrix) LoadBuffer(bits []byte, block int) {
    block &= 0x01

    for i := 0; i < len(bits) ; i++ {
        d.buffer[d.altIndex[i + block * 8]] = bits[i]
    }
}

// A wrapper for WriteBlockData for displaying the buffer.
//
func (d *Adafruit816LedMatrix) DrawBuffer() {
    device := d.ht16k33.Connection()
    if device != nil {
        device.WriteBlockData(0, d.buffer)
    }
}

// Rotates the buffer contents from left to right.
//
func (d *Adafruit816LedMatrix) RotateBuffer() {
    end := d.buffer[d.altIndex[len(d.buffer) - 1]]

    for i := len(d.buffer) - 1 ; i > 0 ; i-- {
        d.buffer[d.altIndex[i]] = d.buffer[d.altIndex[i-1]]
    }

    d.buffer[0] = end
}

