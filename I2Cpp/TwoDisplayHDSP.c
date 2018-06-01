/*
 * CheckHDSP.c:
 * HDSP 211x test code
 *
 * Copyright (c) 2017 William H. Beebe, Jr.
 */

#include <iostream>
#include <csignal>

#include <wiringPi.h>
#include <mcp23017.h>

#include "HDSP.h"

/*
 * This signal handler handles all signals, but ^C or SIGINT
 * in particular. Thus, hitting ^C will reset and turn off the
 * HDSP displays before exiting in an orderly fashion.
 */

void signalHandler(int signal) {
    resetHDSP();
    std::cout << std::endl;
    exit(0);
}

/*
 * Display NULLs across the HDSP display. Should produce leftward
 * pointing arrowheads.
 */

void runLeft() {
    for (int i = MAX_DIGITS; i >= 0; --i) {
        writeCharacter(i, 0); 
        delay(25);
    }
    for (int i = MAX_DIGITS; i >= 0; --i) {
        writeCharacter(i, ' ');
        delay(25);
    }
}

int main (int argc, char *argv[]) {
    signal(SIGINT, signalHandler);

    setupAndReset();
    for (int i = 0; i < 9; ++i) { runLeft(); }
    resetHDSP();

    return 0 ;
}

