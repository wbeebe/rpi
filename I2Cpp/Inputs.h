/*
 * Inputs.h:
 * Input defines for an MCP23017 at I2C address 0x21.
 *
 * Copyright (c) 2017 William H. Beebe, Jr.
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
