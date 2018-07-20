#!/bin/bash
#
# A simple clock that uses the Bash shell to send the Unix date and time to a pair of
# Adafruit displays. CTRL C is trapped to turn off the displays before exiting from
# the script.
#
trap ctrl_c INT

function ctrl_c() {
    display clear > /dev/null
    exit
}
#
# Get the date from the Unix date command, using +%r to get the time formatted in the
# locale's time format. This will give it in 12 hour format, with an AM or PM.
# Then use the Bash shell's built-in regex to replace the last ':' and second digits
# with a space, so that the time string is in the format '12:00 PM' and will print
# across the eight Adafruit hexadecimal digits.
#
while true; do
    DATE=$(date +%r)
    display print "${DATE//:[0-9][0-9] / }" > /dev/null;
    sleep 1;
done

