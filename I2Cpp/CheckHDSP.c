/*
 * Copyright (c) 2018 William H. Beebe, Jr.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * CheckHDSP.c:
 * HDSP 211x test code
 */

#include <iostream>
#include <csignal>

#include <wiringPi.h>
#include <mcp23017.h>

#include "HDSP.h"

/*
 * This signal handler handles all signals, but ^C or SIGINT
 * in particular. Thus, hitting ^C will reset and turn off the
 * HDSP displays before exiting in an orderly fashion.
 */

void signalHandler(int signal) {
    resetHDSP();
    std::cout << std::endl;
    exit(0);
}

const char TEST_MESSAGE[] =
    "                "
    "\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f"
    "\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f "
    "!\"#$%&\'()*+,-./ "
    "0123456789 "
    "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
    "abcdefghijklmnopqrstuvwxyz "
    "{|}~\xff"
    "                ";

/*
 * Scroll a test message across the display.
 * This displays the lower 127 ASCII characters except NULL.
 */

int main (int argc, char *argv[]) {
    signal(SIGINT, signalHandler);

    setupAndReset();
    scrollText(TEST_MESSAGE, 400);
    resetHDSP();

    return 0 ;
}

