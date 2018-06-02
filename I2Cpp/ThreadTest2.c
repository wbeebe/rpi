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
 * ThreadTest.c:
 * 
 * Combines the date clock with checking inputs via threading.
 */

#include <iostream>
#include <csignal>
#include <cstdbool>
#include <thread>

#include <wiringPi.h>
#include <mcp23017.h>

#include "HDSP.h"
#include "Inputs.h"

/*
 * This signal handler handles all signals, but ^C or SIGINT
 * in particular. Thus, hitting ^C will reset and turn off the
 * HDSP displays before exiting in an orderly fashion.
 */
bool    canExecute = TRUE;
bool    canDisplay = TRUE;
bool    isDisplaying = FALSE;

void signalHandler(int signal) {
    canExecute = FALSE;
}

void threadedTime() {
    while (canExecute) {

        if (canDisplay) {
            isDisplaying = TRUE;
            doClock();
            doDate();
            isDisplaying = FALSE;
            delay(100);
        }

        isDisplaying = FALSE;
        delay(100);
    }
}

void overwriteTimeDisplay(const char *text) {
    canDisplay = FALSE;
    while (isDisplaying) { delay(5); }
    writeText(text);
}

/*
 * For testing purposes, or to just make it run while working on the
 * Raspberry Pi, start this program and put it in the background.
 * Bring it back into the foreground and type ^C to turn off display
 * and exit.
 */

int main (int argc, char *argv[]) {
    signal(SIGINT, signalHandler);

    setupAndReset();
    setupInputs();

    std::thread ttime(&threadedTime);
    ttime.detach();

    while( canExecute ) {
        if (!digitalRead(INPT_GPB0)) { overwriteTimeDisplay("GPB0    "); }
        else if (!digitalRead(INPT_GPB1)) { overwriteTimeDisplay("GPB1    "); }
        else if (!digitalRead(INPT_GPB2)) { overwriteTimeDisplay("GPB2    "); }
        else if (!digitalRead(INPT_GPB3)) { overwriteTimeDisplay("GPB3    "); }
        else if (!digitalRead(INPT_GPB4)) { overwriteTimeDisplay("GPB4    "); }
        else if (!digitalRead(INPT_GPB5)) { overwriteTimeDisplay("GPB5    "); }
        else if (!digitalRead(INPT_GPB6)) { overwriteTimeDisplay("GPB6    "); }
        else if (!digitalRead(INPT_GPB7)) { overwriteTimeDisplay("GPB7    "); }
        else { if (!canDisplay) { resetHDSP(); } canDisplay = TRUE; }

        delay(10);
    }

    resetHDSP();
    std::cout << std::endl;
    return 0 ;
}

