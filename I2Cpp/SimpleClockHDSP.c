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
 * SimpleClockHDSP.c:
 * A very simple HDSP 211x clock example.
 * Only uses left-most HDSP display.
 */

#include <iostream>
#include <csignal>
#include <cstdbool>

#include <wiringPi.h>
#include <mcp23017.h>

#include "HDSP.h"

/*
 * This signal handler handles all signals, but ^C or SIGINT
 * in particular. Thus, hitting ^C will reset and turn off the
 * HDSP displays before exiting in an orderly fashion.
 */
bool    canExecute = TRUE;

void signalHandler(int signal) {
    canExecute = FALSE;
}

/*
 * For testing purposes, or to just make it run while working on the
 * Raspberry Pi, start program and put it in the background.
 * Bring it back into the foreground and type ^C to turn off display
 * and exit.
 */

int main (int argc, char *argv[]) {
    signal(SIGINT, signalHandler);

    setupAndReset();

    while (canExecute) {
        doClock( 1 );
        delay(950);
    }

    resetHDSP();
    std::cout << std::endl;
    return 0 ;
}

