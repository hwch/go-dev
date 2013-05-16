package main

import (
        "fmt"
        "net"
)

func main() {
        ipc, err := net.ListenIP("ip4:udp", &net.IPAddr{IP: net.ParseIP("127.0.0.1")})
        if err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }
        bf := make([]byte, 8192)

        for {
                if n, err := ipc.Read(bf); err != nil {
                        fmt.Printf("Error: %v\n", err)
                        return
                } else {
                        fmt.Printf("Recieve[127.0.0.1]:%v\n", bf[:n])
                }
        }
}
