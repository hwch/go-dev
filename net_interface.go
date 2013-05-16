package main

import "fmt"
import "net"

func main() {
        x, err := net.InterfaceAddrs()
        if err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }
        for _, v := range x {
                fmt.Printf("%v\n", v.String())
        }
}
