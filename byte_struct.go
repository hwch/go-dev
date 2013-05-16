package main

import (
        "fmt"
        "unsafe"
)

type UnByteStruct struct {
        u64     uint64
        u16     uint16
        u32     uint32
}

func main() {
        ubs := new(UnByteStruct)

        ubs.u16 = 0x1122
        ubs.u32 = 0x33445566
        ubs.u64 = 0xfffffffffffff

        p1 := (*[16]byte)(unsafe.Pointer(ubs))
        p2 := (*[]byte)(unsafe.Pointer(ubs))

        fmt.Printf("ubs=%p|p1=%p|p2=%p\n", ubs, p1, p2)
        str1 := "p1=0x"
        for i := 0; i < 16; i++ {
                str1 = str1 + fmt.Sprintf("%02x", (*p1)[i])
        }
        fmt.Printf("%s\n", str1)
        fmt.Printf("&p2[0]=0x%x\n", &(*p2)[0])
        fmt.Printf("p2[0]=0x%x\n", (*p2)[0])
}
