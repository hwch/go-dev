package main

import (
        "encoding/gob"
        "fmt"
        "os"
)

func main() {
        var g_Unicode2Gbk map[uint64]uint64
        var g_Gbk2Unicode map[uint64]uint64
        var k uint64

        g_Unicode2Gbk = make(map[uint64]uint64)
        g_Gbk2Unicode = make(map[uint64]uint64)

        rf, err := os.Open("data.db")
        if err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }
        defer rf.Close()
        gb := gob.NewDecoder(rf)
        if err := gb.Decode(&g_Gbk2Unicode); err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }
        if err := gb.Decode(&g_Unicode2Gbk); err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }
        k = 0x57bc
        if v, ok := g_Unicode2Gbk[k]; ok {
                fmt.Printf("0x%x=0x%x\n", k, v)
                if v0, ok0 := g_Gbk2Unicode[v]; ok0 {
                        fmt.Printf("0x%x=0x%x\n", v, v0)
                } else {
                        fmt.Printf("Not found g_Gbk2Unicode\n")
                }
        } else {
                fmt.Printf("Not found\n")
        }
}
