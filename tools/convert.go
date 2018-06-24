/*
Copyright (c) 2018 William H. Beebe, Jr.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

// Reads in a simple text file that contains the binary encoded
// values for the Adafruit Quad Alphanumeric Display and converts
// them to hexadecimal values in Go language.

func main() {
    data, _ := ioutil.ReadFile("table.txt")
    str := string(data)
    lines := strings.Split(str, "\n")
    fmt.Printf("var alphaTable = map[string]uint16 {\n")
    for _, line := range lines {
        if len(line) > 0 {
            segs := strings.Split(line, ", ")
            val, _ := strconv.ParseUint(segs[0], 2, 16)
            fmt.Printf("    %s:0x%04X,\n", segs[1], val)
        }
    }
    fmt.Printf("}\n")
}
