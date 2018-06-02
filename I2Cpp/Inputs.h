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
 * Inputs.h:
 * Input defines for an MCP23017 at I2C address 0x21.
 */

#ifndef INPUTS_H
#define INPUTS_H

#define MCP23017_2  0x21

// Defines for the MCP23017 #2 port expander.

#define	INPT_BASE   120

#define INPT_GPA0   INPT_BASE
#define INPT_GPA1   INPT_BASE+1
#define INPT_GPA2   INPT_BASE+2
#define INPT_GPA3   INPT_BASE+3
#define INPT_GPA4   INPT_BASE+4
#define INPT_GPA5   INPT_BASE+5
#define INPT_GPA6   INPT_BASE+6
#define INPT_GPA7   INPT_BASE+7

#define INPT_GPB0   INPT_BASE+8
#define INPT_GPB1   INPT_BASE+9
#define INPT_GPB2   INPT_BASE+10
#define INPT_GPB3   INPT_BASE+11
#define INPT_GPB4   INPT_BASE+12
#define INPT_GPB5   INPT_BASE+13
#define INPT_GPB6   INPT_BASE+14
#define INPT_GPB7   INPT_BASE+15

void setupInputs();

#endif
