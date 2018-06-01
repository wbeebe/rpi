/*
 * ResetHDSP.c:
 * Reset HDSP 211x to turn off display digits.
 *
 * Copyright (c) 2017 William H. Beebe, Jr.
 */

#include <wiringPi.h>
#include <mcp23017.h>

#include "HDSP.h"

int main (int argc, char *argv[]) {
    setupAndReset();
    return 0 ;
}

