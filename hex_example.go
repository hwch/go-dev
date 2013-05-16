package main

import (
        "encoding/hex"
        "fmt"
)

func main() {
        s := "feca31234abcdef375823951"
        if x, err := hex.DecodeString(s); err != nil {
                fmt.Printf("Error:%v\n", err)
        } else {
                fmt.Printf("%v\n", x)
        }
        x := make([]byte, 4)
        x[0] = 0xff
        x[1] = 0xab
        x[2] = 0x1f
        x[3] = 0xf0
        fmt.Printf("%s\n", hex.EncodeToString(x))
}
