/*
 * Copyright (c) 2018 William H. Beebe, Jr.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * ScrollText.c:
 * HDSP 211x scroll a line of text across the display.
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

