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
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"gobot.io/x/gobot/drivers/i2c"

	"github.com/wbeebe/rpi/devices"
)

// DefaultAddress is the default I2C address at which the device can be
// addressed. It is in fact the lowest address any of these devices can
// be addressed.
//
const DefaultAddress int = 0x70

func lightAll(device i2c.Connection) {
	// First four digits for Alphanumeric and 8x16 Matrix
	// FeatherWing Displays.
	//
	// The 'digit' address is the address/offset into the
	// HT16K33's internal eight byte array. Each bit
	// represents a segment or LED, each address a section
	// within an entire device
	//
	// Digit 0
	//
	device.WriteWordData(0, 0xFFFF)

	// Digit 1
	//
	device.WriteWordData(2, 0xFFFF)

	// Digit 2
	//
	device.WriteWordData(4, 0xFFFF)

	// Digit 3
	//
	device.WriteWordData(6, 0xFFFF)

	// Rest of the bytes for the
	// Adafruit 0.8" 8x16 LED Matrix FeatherWing Display
	device.WriteWordData(8, 0xFFFF)
	device.WriteWordData(10, 0xFFFF)
	device.WriteWordData(12, 0xFFFF)
	device.WriteWordData(14, 0xFFFF)
}

func main() {
	// Hook the various system abort calls for us to use or ignore as we
	// see fit. In particular hook SIGINT, or CTRL+C for below.
	//
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	// We can pass zero to many addresses on the command line.
	// For example, 'raw_ht16k33 0x70 0x71' will turn on both
	// displays, one after the other, if they are both attached.
	// If either one is not there or unreachable, the application
	// will abort.
	//
	var addresses []int

	for _, arg := range os.Args[1:] {
		if newAddress, err := strconv.ParseInt(arg, 0, 32); err == nil {
			addresses = append(addresses, int(newAddress))
		} else {
			log.Fatal(err)
		}
	}

	// If nothing passed on the command line then use the default
	// address/
	//
	if len(addresses) == 0 {
		addresses = append(addresses, DefaultAddress)
	}

	// We want to capture CTRL+C to first clear the display and then exit.
	// We don't want to leave the display lit on an abort.
	//
	ht := devices.NewHT16K33Driver(DefaultAddress)

	go func() {
		for {
			signal := <-signalChan
			switch signal {
			case syscall.SIGINT:
				// CTRL+C
				fmt.Println()
				ht.Clear()
				ht.Close()
				os.Exit(0)
			default:
			}
		}
	}()

	// Iterate over all the addresses passed on the command line (or not),
	// aborting if any of the devices are unreachable.
	//
	for _, addr := range addresses {
		address := int(addr)
		ht16k33 := devices.NewHT16K33Driver(address)
		ht = ht16k33

		err := ht16k33.Start()
		if err != nil {
			log.Fatal(err)
		}

		ht16k33.Clear()
		lightAll(ht16k33.Connection())
		time.Sleep(2 * time.Second)
		ht16k33.Clear()
		ht16k33.Close()
	}
}
