package main

import "fmt"
import "net"

func main() {
        uc, err := net.DialUDP("udp4",
                nil,
                &net.UDPAddr{IP: net.IPv4bcast, Port: 2425})
        if err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }
        defer uc.Close()
        if _, err := uc.Write([]byte("Hello World")); err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }

}
