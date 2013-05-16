package main

import (
        "fmt"
        "time"
)

func main() {
        c := make(chan int, 4)
        d := make(chan bool)
        defer close(c)
        go func(chan int) {
                defer func() {
                        d <- true
                }()
                i := 10
                for {
                        c <- i
                        time.Sleep(time.Second)
                        i++
                }
        }(c)

        go func(chan int) {
                defer func() {
                        d <- true
                }()
                i := 0
                for {
                        x := <-c
                        fmt.Printf("%d|%d\n", i, x)
                        i++
                }
        }(c)
        <-d
        <-d
}
