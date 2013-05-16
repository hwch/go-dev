package main

import (
        "fmt"
        "strconv"
        "time"
)

func thread_func(c chan string, v int) {
        time.Sleep(time.Second * time.Duration(v))
        c <- "This is the " + strconv.Itoa(v) + " times"
}

func main() {
        c := make(chan string)
        f := func(c chan string) {
                close(c)
                fmt.Println("channel is closed")
        }
        defer f(c)
        for i := 0; i < 12; i++ {
                go thread_func(c, i)
        }

        i := 0
L:
        for {
                if i == 12 {
                        break L
                }
                v, ok := <-c
                if ok {
                        fmt.Printf("Recieve => %v\n", v)
                        i++
                }
                fmt.Printf("Nothing\n")
        }
}
