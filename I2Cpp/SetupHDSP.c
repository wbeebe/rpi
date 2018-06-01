/*
 * HDSP2111.c:
 * Setup the MCP23017 to HDSP 211x display.
 *
 * Copyright (c) 2017 William H. Beebe, Jr.
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

