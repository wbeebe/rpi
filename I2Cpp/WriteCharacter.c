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
 * WriteCharacter.c:
 * A very simple HDSP 211x write a character using MCP23017.
 */
#include <stdio.h>

#include <wiringPi.h>
#include <mcp23017.h>

#include "HDSP.h"

/*
 * Write a single character to the display at a given character address.
 * Character addresses range from 0 (far left) to 7 (far right) inclusive.
 */

void writeCharacter(const size_t addr, const int character) {
    if ((addr < 0) || (addr > MAX_DIGITS)) {
        return;
    }

    digitalWrite(HDSP_CE1, HIGH);
    digitalWrite(HDSP_CE2, HIGH);

    // Pick a character address using A0 to A2
    //
    digitalWrite(HDSP_A0, addr & 0x01);
    digitalWrite(HDSP_A1, addr & 0x02);
    digitalWrite(HDSP_A2, addr & 0x04);

    // Write out the lowest seven bits of the character.
    // Bit 8 isn't used.
    //
    digitalWrite(HDSP_D0, character & 0x01);
    digitalWrite(HDSP_D1, character & 0x02);
    digitalWrite(HDSP_D2, character & 0x04);
    digitalWrite(HDSP_D3, character & 0x08);
    digitalWrite(HDSP_D4, character & 0x10);
    digitalWrite(HDSP_D5, character & 0x20);
    digitalWrite(HDSP_D6, character & 0x40);

    // Write out the address and data setup above.
    // Decide which Display, 0 (left-most) or 1 (right_most)
    //
    if (addr & 0x08) {
        // Right-most display
        //
        digitalWrite(HDSP_CE2, LOW);
        digitalWrite(HDSP_WR2, LOW);

        digitalWrite(HDSP_WR2, HIGH);
        digitalWrite(HDSP_CE2, HIGH);
    }
    else {
        // Left-most display
        //
        digitalWrite(HDSP_CE1, LOW);
        digitalWrite(HDSP_WR1, LOW);

        digitalWrite(HDSP_WR1, HIGH);
        digitalWrite(HDSP_CE1, HIGH);
    }
}

