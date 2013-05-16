package main

import (
        "fmt"
        "net"
)

type myByte []byte

func (b myByte) String() string {
        s := ""
        for i := 0; i < len(b); i++ {
                s = s + fmt.Sprintf("[%08b]", b[i])
        }
        return s
}

func main() {
        //ripa, err := net.ResolveIPAddr("ip", "192.168.135.253")
        ripa, err := net.ResolveIPAddr("ip", "localhost")
        if err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }
        lp, err1 := net.ListenIP("ip:icmp", ripa)
        if err1 != nil {
                fmt.Printf("Error1: %v\n", err1)
                return
        }
        for {
                buffer := make(myByte, 4096)
                if n, addr, err := lp.ReadFrom(buffer); err != nil {
                        fmt.Printf("Error2: %v\n", err)
                        return
                } else {
                        fmt.Printf("Revieve from %v => %v\n", addr, buffer[:n])
                }

        }
}
