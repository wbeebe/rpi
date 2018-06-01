/*
 * WriteText.c:
 * HDSP 211x write a single line of text across the display.
 *
 * Copyright (c) 2017 William H. Beebe, Jr.
 */

#include <string.h>

#include <wiringPi.h>
#include <mcp23017.h>

#include "HDSP.h"

void writeText(const char *text) {
    size_t length = strlen(text);
    if (length <= MAX_DIGITS) {
        for (size_t i = 0; i < length; ++i) {
            writeCharacter(i, text[i]);
        }
    }
}

