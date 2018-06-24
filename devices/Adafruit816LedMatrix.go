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
    "gobot.io/x/gobot/drivers/i2c"
)

var buffer []byte = make([]byte, 16)
var altIndex []int = []int{0,2,4,6,8,10,12,14,1,3,5,7,9,11,13,15}

// Loads the buffer with data, in the pattern necessary for proper
// displaying. Works with the concept of blocks that matches the
// 8x8 LED arrays on the display. Block 0 is on the left, block 1
// on the right.
//
func LoadBuffer(bits []byte, block int) {
    block &= 0x01

    for i := 0; i < len(bits) ; i++ {
        buffer[altIndex[i + block * 8]] = bits[i]
    }
}

// A wrapper for WriteBlockData for displaying the buffer.
//
func DrawBuffer(device i2c.Connection) {
    device.WriteBlockData(0, buffer)
}

// Rotates the buffer contents from left to right.
//
func RotateBuffer() {
    end := buffer[altIndex[len(buffer) - 1]]

    for i := len(buffer) - 1 ; i > 0 ; i-- {
        buffer[altIndex[i]] = buffer[altIndex[i-1]]
    }

    buffer[0] = end
}

// Turns every lit LED off by writing binary zeros to all locations.
//
func ClearAll816(device i2c.Connection) {
    block := make([]byte, 16)
    device.WriteBlockData(0, block)
}

