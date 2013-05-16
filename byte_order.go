package main

import (
        "fmt"
        "unsafe"
)

func isLittleEndian() bool {
        var x int32 = 0x12345678
        p := unsafe.Pointer(&x)
        p1 := (*[4]byte)(p)
        return p1[0] == 0x78
}

func main() {
        fmt.Printf("sizeof(int)=%d\n", unsafe.Sizeof(int(1)))
        if !isLittleEndian() {
                fmt.Println("本机器：大端")
        } else {
                fmt.Println("本机器：小端")
        }
}
