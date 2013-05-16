package main

import (
        "fmt"
        "unsafe"
)

func main() {
        x := []byte{0x1, 0x2, 0x3, 0x4, 0x5}
        fmt.Printf("%d\n", unsafe.Sizeof(x))
}
