package main

import (
        "flag"
        "fmt"
        "net"
        "runtime"
        "sync"
)

func main() {
        var ip string
        var port int
        var max_thread int
        flag.StringVar(&ip, "ip", "", "Other IP address")
        flag.IntVar(&port, "port", -1, "The port that you want to scan")
        flag.IntVar(&max_thread, "n", 1, "The number process of scan target")
        flag.Parse()

        runtime.GOMAXPROCS(runtime.NumCPU())
        if ip == "" {
                fmt.Print("Please input ip address\n")
                flag.PrintDefaults()
                return
        }
        if max_thread <= 1 {
                max_thread = 1
        }
        if port <= 0 {
                unit := 0xffff / max_thread
                mod := 0xffff % max_thread
                count := 1
                wg := new(sync.WaitGroup)

                for i := 0; i < max_thread; i++ {
                        wg.Add(1)
                        go func(c, u int) {
                                ip1 := ip
                                defer wg.Done()
                                for j := c; j < c+u; j++ {
                                        host_port := fmt.Sprintf("%s:%d", ip1, j)
                                        if c, err := net.Dial("tcp", host_port); err != nil {
                                                // fmt.Printf("IP[%s]Port[%d] => Error: %v\n", ip, j, err)
                                        } else {
                                                fmt.Printf("#####IP[%s]Port[%d] opened #####\n", ip1, j)
                                                c.Close()
                                        }
                                }
                        }(count, unit)
                        count += unit
                }
                if mod != 0 {
                        for ; count < 0x10000; count++ {
                                host_port := fmt.Sprintf("%s:%d", ip, count)
                                if c, err := net.Dial("tcp", host_port); err != nil {
                                        //fmt.Printf("IP[%s]Port[%d] => Error: %v\n", ip, count, err)
                                } else {
                                        fmt.Printf("#####IP[%s]Port[%d] opened #####\n", ip, count)
                                        c.Close()
                                }
                        }
                }
                wg.Wait()

        } else {
                host_port := fmt.Sprintf("%s:%d", ip, port)
                if c, err := net.Dial("tcp", host_port); err != nil {
                        fmt.Printf("IP[%s]Port[%d] => Error: %v\n", ip, port, err)
                } else {
                        c.Close()
                }
        }

}
