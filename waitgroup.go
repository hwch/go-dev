package main

import (
        "fmt"
        "sync"
        "time"
)

func main() {
        wg := new(sync.WaitGroup)

        for i := 0; i < 5; i++ {
                wg.Add(1)
                if i == 3 {
                        go func(d int) {
                                defer wg.Done()
                                time.Sleep(time.Second * 3)
                                fmt.Printf("This is %dth goroutine\n", d)
                                panic("This is test")
                        }(i)
                } else {
                        go func(d int) {
                                defer wg.Done()
                                time.Sleep(time.Second * 2)
                                fmt.Printf("This is %dth goroutine\n", d)
                        }(i)
                }
        }

        wg.Wait()
}
