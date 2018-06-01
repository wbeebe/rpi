/*
 * DateModule.c:
 * A very simple HDSP 211x date example.
 *
 * Copyright (c) 2017 William H. Beebe, Jr.
 */

#include <chrono>

#include <wiringPi.h>
#include <mcp23017.h>

#include "HDSP.h"

const char *months[] = {
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Dec"
};

void doDate() {
    // Hard coded to adjust for DST on the east coast.
    // TODO: Make this not hard coded.
    //
    std::chrono::time_point<std::chrono::system_clock> utc = std::chrono::system_clock::now();
    std::chrono::time_point<std::chrono::system_clock> now = utc - std::chrono::hours(4);

    time_t current = std::chrono::system_clock::to_time_t(now);
    const struct tm *ptr_time = gmtime(&current);

    const char *month = months[ptr_time->tm_mon];
    writeCharacter(8, month[0]);
    writeCharacter(9, month[1]);
    writeCharacter(10, month[2]);
    writeCharacter(11, ' ');
    writeCharacter(12, ptr_time->tm_mday > 9 ? ('0' + ptr_time->tm_mday/10) : ' ');
    writeCharacter(13, '0' + ptr_time->tm_mday%10);

    if (ptr_time->tm_mday == 1 || ptr_time->tm_mday == 21 || ptr_time->tm_mday == 31) {
        writeCharacter(14, 's');
        writeCharacter(15, 't');
    }
    else if (ptr_time->tm_mday == 2 || ptr_time->tm_mday == 22) {
        writeCharacter(14, 'n');
        writeCharacter(15, 'd');
    }
    else if (ptr_time->tm_mday == 3 || ptr_time->tm_mday == 23) {
        writeCharacter(14, 'r');
        writeCharacter(15, 'd');
    }
    else {
        writeCharacter(14, 't');
        writeCharacter(15, 'h');
    }
}

