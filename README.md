# rpi

A collection of tools to manipulate the Raspberry Pi GPIO and I2C bus. Some tools are written in Go, and some in C++

## Go Apps

The Go apps were started with Go 1.10.3 and use the Gobot framework (https://gobot.io). The most recent version of Go used to build the apps is 1.15.3.

The Go software is downloaded and installed from https://golang.org/ . Do not install Go from any Linux repo unless it matches the release version on the website.

## C++ Apps (Deprecated)

The C++ apps were initially built with GCC 6.3 and Gordon Henderson's excellent wiringPi framework (http://wiringpi.com/). However, as of August 2018, Gordon Henderson has deprecated his framework and is no longer developing nor supporting it. Therefore, these apps are deprecated as well.

## Hardware

Everything has been built and tested on a Raspberry Pi 3 Model B+, Raspberry Pi 3 Model B, a Raspberry Pi Zero W, and Raspberry Pi 4B (1GiB, 2GiB, 4GiB and 8GiB)

## Operating System

The OS on all devices is the latest 64-bit version of Raspbian available from https://downloads.raspberrypi.org/raspios_arm64/images/ . Full blown Raspbian (with graphical desktop) is installed on the
Raspberry Pi 4Bs. Raspbian Lite is on the Zero.

## Data Sheets

It's hard finding data sheets for electronic components these days. To that end I've collected as many as possible for every component I use in my projects and put them in the datasheets folder. I hope you find them useful in your own work.

The folder contains the following data sheets in PDF format:

+ AV02-0629EN-DS-HDSP-210x-02Dec20100.pdf - 
 A current datasheet for the HP HDSP-2111 and HDSP-2113 which I still
have a number of. Acquired back in the early 1980s, these are eight
character alphanumeric displays. They are driven through MCP23017s
in one of my current designs. Look in the I2Cpp folder for the
software.

+ BCM2837-ARM-Peripherals-Stanford-EDU-Corrections.pdf - 
 The BCM2835 datasheet with corrections for the BCM2837 applied by
Stanford University's Computer Science division for classroom work.

+ BST_BNO055_DS000_14.pdf - 
 The Bosch BNO055 Sensortec 9-axis motion sensor. This part is used
in the Adafruit Absolute Orientation IMU Fusion Breakout

+ DL2416.pdf - 
 The DL2416 is a small four digit alphanumeric display. This was
a big seller in the 1980s.

+ HP_5082-7340.pdf - 
 A single digit alphanumeric display. HP back in the 1980s had
a semiconductor and test instrument division. These displays were
made by them. They were both used internally in their test equipment
as well as sold to other customers.

+ HT16K33V110.pdf - 
 The HT16K33 is the I2C peripheral chip behind many of Adafruit's
FeatherWing displays.

+ MCP23017.pdf - 
 An 8-bit dual port I2C port expander. Complete pinouts and
I2C command and control register layout.

+ UM10204_I2C_BUS_SPEC_AND_USER_MANUAL.pdf - 
 The I2C Bus Specification and User Manual. Use this to better
understand what the hardware and drivers are (attempting to) doing
on the Raspberry Pi.

## License
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

