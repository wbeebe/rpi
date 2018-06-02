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

/*
 * Display NULLs across the HDSP display. Should produce leftward
 * pointing arrowheads.
 */

void runLeft() {
    for (int i = 7; i >= 0; --i) {
        writeCharacter(i, 0); 
        delay(100);
    }
    for (int i = 7; i >= 0; --i) {
        writeCharacter(i, ' ');
        delay(100);
    }
}

int main (int argc, char *argv[]) {
    signal(SIGINT, signalHandler);
    setupAndReset();
    for (int i = 0; i < 5; ++i) { runLeft(); }
    resetHDSP();
    return 0 ;
}

