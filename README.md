# rpi

A collection of tools to manipulate the Raspberry Pi GPIO and I2C bus.
Some tools are written in Go, and some in C++

## Go Tools

The Go tools are built with Go 1.10.2 and use the Gobot framework.
The Go tools were downloaded and installed from https://golang.org/
Do not install Go from any Linux repo unless it matches the release
version on the website.

## C++ Tools

The C++ tools are built with GCC 6.3 and use the wiringPi framework.
C++ is provided in the Linux repo and is sufficiently advanced.

## Hardware

Everything has been built and tested on a Raspberry Pi 3 Model B+,
Raspberry Pi 3 Model B, and a Raspberry Pi Zero W.

## Operating System

The OS on all devices is the latest version of Raspbian available
from https://www.raspberrypi.org/downloads/raspbian/
Full blown Raspbian (with graphical desktop) is installed on the
Raspberry Pi 3s. Raspbian Lite is on the Zero.

## Data Sheets

It's hard finding data sheets for electronic components these days.
To that end I've collected as many as possible for every component I
use in my projects and put them in the datasheets folder. I hope you
find them useful in your own work.

The folder contains the following data sheets in PDF format:

BCM2837-ARM-Peripherals-Stanford-EDU-Corrections.pdf
The BCM2835 datasheet with corrections for the BCM2837 applied by
Stanford University's Computer Science division for classroom work.

BST_BNO055_DS000_14.pdf
The Bosch BNO055 Sensortec 9-axis motion sensor. This part is used
in the Adafruit Absolute Orientation IMU Fusion Breakout

DL2416.pdf
The DL2416 is a small four digit alphanumeric display. This was
a big seller in the 1980s.

HP_5082-7340.pdf
A single digit alphanumeric display. HP back in the 1980s had
a semiconductor and test instrument division. These displays were
made by them. They were both used internally in their test equipment
as well as sold to other customers.

HT16K33V110.pdf
The HT16K33 is the I2C peripheral chip behind many of Adafruit's
FeatherWing displays.

MCP23017.pdf
An 8-bit dual port I2C port expander. Complete pinouts and
I2C command and control register layout.

UM10204_I2C_BUS_SPEC_AND_USER_MANUAL.pdf
The I2C Bus Specification and User Manual. Use this to better
understand what the hardware and drivers are (attempting to) doing
on the Raspberry Pi.
