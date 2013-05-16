package main

import (
        "fmt"
        "net"
        "runtime"
)

func HandleConnections(conn net.Conn) {
        var n int
        var err error
        mc := make([]byte, 1024)
        defer conn.Close()
        for {
                if n, err = conn.Read(mc); err != nil {
                        if err.Error() == "EOF" {
                                fmt.Printf("Other closed\n")
                        } else {
                                fmt.Printf("Read error: %v\n", err)
                        }

                        return
                }

                fmt.Printf("Recieve: %s\n", string(mc[:n]))
                if _, err := conn.Write(mc); err != nil {
                        fmt.Printf("Write error: %v\n", err)
                        return
                }
                for i := 0; i < len(mc); i++ {
                        mc[i] = 0x0
                }
        }
}

func main() {
        runtime.GOMAXPROCS(2)
        ln, err := net.Listen("tcp", ":25125")
        if err != nil {
                fmt.Printf("Listen error: %v\n", err) // handle error
                return
        }
        for {
                conn, err := ln.Accept()
                if err != nil {
                        fmt.Printf("Accept error: %v\n", err) // handle error
                        continue
                }
                go HandleConnections(conn)
        }
}
