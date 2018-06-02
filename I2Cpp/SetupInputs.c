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
 * SetupInputs.c:
 * Setup MCP23017 #2 to handle inputs.
 */

#include <wiringPi.h>
#include <mcp23017.h>
#include "Inputs.h"

/*
 * Setup the MCP23017 to write ASCII text to the INPT 2111 display.
 */

void setupInputs() {
    wiringPiSetup();
    mcp23017Setup(INPT_BASE, MCP23017_2);

    pinMode(INPT_GPB0, INPUT);
    pullUpDnControl(INPT_GPB0, PUD_OFF);
    pinMode(INPT_GPB1, INPUT);
    pullUpDnControl(INPT_GPB1, PUD_OFF);
    pinMode(INPT_GPB2, INPUT);
    pullUpDnControl(INPT_GPB2, PUD_OFF);
    pinMode(INPT_GPB3, INPUT);
    pullUpDnControl(INPT_GPB3, PUD_OFF);
    pinMode(INPT_GPB4, INPUT);
    pullUpDnControl(INPT_GPB4, PUD_OFF);
    pinMode(INPT_GPB5, INPUT);
    pullUpDnControl(INPT_GPB5, PUD_OFF);
    pinMode(INPT_GPB6, INPUT);
    pullUpDnControl(INPT_GPB6, PUD_OFF);
    pinMode(INPT_GPB7, INPUT);
    pullUpDnControl(INPT_GPB7, PUD_OFF);
}

