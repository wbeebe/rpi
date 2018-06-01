/*
 * CheckInputs.c:
 * MCP23017 port B input test code
 *
 * Copyright (c) 2017 William H. Beebe, Jr.
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

