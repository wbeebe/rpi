package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

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
