/*
 * SetupInputs.c:
 * Setup MCP23017 #2 to handle inputs.
 *
 * Copyright (c) 2017 William H. Beebe, Jr.
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

