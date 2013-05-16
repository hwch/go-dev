package main

import (
        "fmt"
)

func main() {
        var x uint
        x = 0x1
        str := ""
        switch {
        case x&0x1 != 0:
                str = str + "<1>"
                fallthrough
        case x&0x2 != 0:
                str = str + "<2>"
                fallthrough
        case x&0x4 != 0:
                str = str + "<4>"
                fallthrough
        case x&0x8 != 0:
                str = str + "<8>"
                fallthrough
        case x&0x10 != 0:
                str = str + "<16>"
        }
        fmt.Printf("%s\n", str)
}
