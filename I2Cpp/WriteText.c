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
 * WriteText.c:
 * HDSP 211x write a single line of text across the display.
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

