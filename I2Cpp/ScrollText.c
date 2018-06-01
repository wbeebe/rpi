/*
 * ScrollText.c:
 * HDSP 211x scroll a line of text across the display.
 *
 * Copyright (c) 2017 William H. Beebe, Jr.
 */

#include <string.h>

#include <wiringPi.h>
#include <mcp23017.h>

#include "HDSP.h"

/*
 * This displays the lower 127 ASCII characters except NULL.
 */

void scrollText(const char text[], int delayMillis) {
    size_t length = strlen(text) - MAX_DIGITS;
    for (size_t i = 0; i < length; ++i) {
        for (size_t j = 0; j <= MAX_DIGITS; ++j) {
            writeCharacter(j, text[i+j]);
        }

        delay(delayMillis);
    }
}

