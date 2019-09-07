#!/bin/bash
#
# Copyright (c) 2019 William H. Beebe, Jr.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# ------------------------------------------------------------------------------
#
# A simple clock that uses the Bash shell to send the Unix date and time to a
# pair of Adafruit displays.
#
# CTRL C is trapped to turn off the displays before exiting from the script.
#
trap ctrl_c INT

function ctrl_c() {
    display clear > /dev/null
    echo
    exit
}
#
# Get the date from the Unix date command, using +%r to get the time formatted
# in the locale's time format. This will give it in 12 hour format, with an
# AM or PM. Then use the Bash shell's built-in regex to replace the last ':'
# and second digits with a space, so that the time string is in the format
# '12:00 PM' and will print across the eight Adafruit hexadecimal digits.
#
while true; do
    DATE=$(date +%r)
    display print "${DATE//:[0-9][0-9] / }" > /dev/null
    sleep 3
    DATE=$(date +%D)
    display print "$DATE" > /dev/null
    sleep 3
    #display clear > /dev/null
done

