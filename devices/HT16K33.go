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

    "gobot.io/x/gobot/drivers/i2c"
    "gobot.io/x/gobot/platforms/raspi"
)

// Constants
//
// These constants are origially from Adafruit's GitHub
// repository C++ files at
// https://github.com/adafruit/Adafruit_LED_Backpack/
// Some changes made by me for clarity.
//

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

type HT16K33Driver struct {
    name string
    address int
    connection i2c.Connection
}

func NewHT16K33Driver(addr int) *HT16K33Driver {
    driver := &HT16K33Driver {
        name: "HT16K33",
        address: addr,
    }

    return driver
}

func (driver *HT16K33Driver) Name() string { return driver.name }
func (driver *HT16K33Driver) SetName(newName string ) { driver.name = newName }
func (driver *HT16K33Driver) Connection() i2c.Connection { return driver.connection }

// Initializes and opens a connection to an HT16K33.
// Returns the i2c.Connection on sucess, err on failure.
//
func (driver *HT16K33Driver) Start() (err error) {
    adapter := raspi.NewAdaptor()
    adapter.Connect()
    bus := adapter.GetDefaultBus()

    // Check to see if the device actually is on the I2C buss.
    // If it is then use it, else return an error.
    //
    if device, err := adapter.GetConnection(driver.address, bus) ; err == nil {
        if _, err := device.ReadByte() ; err == nil {
            fmt.Printf(" Using device 0x%x / %d on bus %d\n", driver.address, driver.address, bus)
        } else {
            return fmt.Errorf(" Could not find device 0x%x / %d", driver.address, driver.address)
        }
    }

    driver.connection, _ = adapter.GetConnection(driver.address, bus)
    // Turn on chip's internal oscillator.
    driver.connection.WriteByte(HT16K33_SYSTEM_SETUP | HT16K33_OSCILLATOR_ON)
    // Turn on the display. YOU HAVE TO SEND THIS.
    driver.connection.WriteByte(HT16K33_DISPLAY_SETUP | HT16K33_DISPLAY_ON)
    // Set for maximum LED brightness.
    driver.connection.WriteByte(HT16K33_CMD_BRIGHTNESS | 0x0f)
    return nil
}

// Clear the device of all data, and in the process turn off
// any LEDs that might be on.
//
func (driver *HT16K33Driver) Clear() {
    buffer := make([]byte, 16)
    driver.connection.WriteBlockData(0, buffer)
}

func (driver *HT16K33Driver) Close() {
    driver.connection.Close()
}
