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
 * HDSP2111.c:
 * Setup the MCP23017 to HDSP 211x display.
 */

#include <wiringPi.h>
#include <mcp23017.h>
#include "HDSP.h"

/*
 * Setup the MCP23017 to write ASCII text to the HDSP 2111 display.
 */

void setupAndReset() {
    wiringPiSetup();
    mcp23017Setup(HDSP_BASE, MCP23017_1);

    pinMode(HDSP_D0, OUTPUT);
    pinMode(HDSP_D1, OUTPUT);
    pinMode(HDSP_D2, OUTPUT);
    pinMode(HDSP_D3, OUTPUT);
    pinMode(HDSP_D4, OUTPUT);
    pinMode(HDSP_D5, OUTPUT);
    pinMode(HDSP_D6, OUTPUT);

    pinMode(HDSP_A0, OUTPUT);
    pinMode(HDSP_A1, OUTPUT);
    pinMode(HDSP_A2, OUTPUT);

    pinMode(HDSP_CE1, OUTPUT);
    pinMode(HDSP_CE2, OUTPUT);
    pinMode(HDSP_WR1, OUTPUT);
    pinMode(HDSP_WR2, OUTPUT);
    pinMode(HDSP_RST, OUTPUT);

    digitalWrite(HDSP_CE1, HIGH);
    digitalWrite(HDSP_CE2, HIGH);
    digitalWrite(HDSP_WR1, HIGH);
    digitalWrite(HDSP_WR2, HIGH);

    resetHDSP();
}

/*
 * Reset the HDSP. Must have run setupAndReset first!
 */
void resetHDSP() {
    digitalWrite(HDSP_RST, LOW);
    digitalWrite(HDSP_RST, HIGH);
}

