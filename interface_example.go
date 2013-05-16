package main

import (
        "fmt"
        "net"
)

func main() {
        var x interface{} = 1
        var y net.TCPConn
        var z net.Conn = &y
        z.Close()
        fmt.Printf("%v\n", x)
}
