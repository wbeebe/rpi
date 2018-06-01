/*
 * ClockModuleHDSP.c:
 * A very simple HDSP 211x clock example.
 *
 * Copyright (c) 2017 William H. Beebe, Jr.
 */

#include <chrono>

#include <wiringPi.h>
#include <mcp23017.h>

#include "HDSP.h"

/*
 * A simple time to display conversion utility.
 * The colon characters are displayed on odd seconds,
 * resulting in a two second flash rate:
 * odd second on, even second off.
 */

void doClock( const bool toggle ) {
    std::time_t current = std::time(0);
    const struct tm *ptr_time = std::localtime(&current);

    writeCharacter(0, '0' + ptr_time->tm_hour/10);
    writeCharacter(1, '0' + ptr_time->tm_hour%10);
    writeCharacter(3, '0' + ptr_time->tm_min/10);
    writeCharacter(4, '0' + ptr_time->tm_min%10);
    writeCharacter(6, '0' + ptr_time->tm_sec/10);
    writeCharacter(7, '0' + ptr_time->tm_sec%10);

    if (!toggle) {
        writeCharacter(2, ':');
        writeCharacter(5, ':');
    }
    else {
        writeCharacter(2, ptr_time->tm_sec % 2 ? ':' : '.');
        writeCharacter(5, ptr_time->tm_sec % 2 ? '.' : ':');
    }
}

