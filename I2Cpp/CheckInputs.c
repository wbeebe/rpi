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
 * CheckInputs.c:
 * MCP23017 port B input test code
 */

#include <iostream>
#include <csignal>

#include <wiringPi.h>
#include <mcp23017.h>

#include "HDSP.h"
#include "Inputs.h"

void signalHandler(int signal) {
    std::cout << std::endl;
    resetHDSP();
    exit(0);
}

int main (int argc, char *argv[]) {
    signal(SIGINT, signalHandler);

    setupAndReset();
    setupInputs();

    while( true ) {
        if (!digitalRead(INPT_GPB0)) { writeText("GPB0    "); }
        else if (!digitalRead(INPT_GPB1)) { writeText("GPB1    "); }
        else if (!digitalRead(INPT_GPB2)) { writeText("GPB2    "); }
        else if (!digitalRead(INPT_GPB3)) { writeText("GPB3    "); }
        else if (!digitalRead(INPT_GPB4)) { writeText("GPB4    "); }
        else if (!digitalRead(INPT_GPB5)) { writeText("GPB5    "); }
        else if (!digitalRead(INPT_GPB6)) { writeText("GPB6    "); }
        else if (!digitalRead(INPT_GPB7)) { writeText("GPB7    "); }
        else { writeText("No input"); }

        delay(100);
    }

    return 0 ;
}

