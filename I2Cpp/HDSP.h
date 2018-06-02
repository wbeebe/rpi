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
 * HDSP2111.h:
 * HDSP 211x defines.
 */

#ifndef HDSP_H
#define HDSP_H
#include <cstring>

#define MCP23017_1  0x20
#define MCP23017_2  0x21

// Defines for the MCP23017 port expander to HDSP2111 display.

#define	HDSP_BASE   100
#define HDSP_D0     HDSP_BASE       // to D0
#define HDSP_D1     HDSP_BASE+1     // to D1
#define HDSP_D2     HDSP_BASE+2     // to D2
#define HDSP_D3     HDSP_BASE+3     // to D3
#define HDSP_D4     HDSP_BASE+4     // to D4
#define HDSP_D5     HDSP_BASE+5     // to D5
#define HDSP_D6     HDSP_BASE+6     // to D6
#define HDSP_D7     HDSP_BASE+7     // to D7

#define HDSP_A0     HDSP_BASE+8     // to A0
#define HDSP_A1     HDSP_BASE+9     // to A1
#define HDSP_A2     HDSP_BASE+10    // to A2

#define HDSP_CE1    HDSP_BASE+11    // Enable Display 1

// This pin is toggled low, then high, to write a charater
// to the HDSL 2111 display.
//
#define HDSP_WR1    HDSP_BASE+12    // Write to Display 1

#define HDSP_CE2    HDSP_BASE+13    // Enable Display 2
#define HDSP_WR2    HDSP_BASE+14    // Write to Display 2
#define HDSP_RST    HDSP_BASE+15    // Reset line to Displays

static const size_t MAX_DIGITS = 15;

void setupAndReset();
void resetHDSP();
void writeCharacter(const size_t addr, const int character);
void writeText(const char *text);
void scrollText(const char text[], int delayMillis);
void doClock( const bool toggle=false );
void doDate();

#endif
