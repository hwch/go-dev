package main

import (
        "fmt"
        "runtime"
        "time"
)

var result chan int

func main() {

        runtime.GOMAXPROCS(runtime.NumCPU())

        go func() {
                result = make(chan int)
                defer close(result)
                sum := 0
                println("This is child process !!!")

                for i := 0; i < 1000; i++ {
                        sum += i
                }
                // time.Sleep(2 * time.Second)
                result <- sum
        }()
        //time.Sleep(2 * time.Second)
        // println("This is parent process !!!")
        for {
                time.Sleep(time.Second)
                if r, ok := <-result; ok == true {
                        fmt.Printf("The result is %d\n", r)
                        break
                } else {
                        println("Not ready")
                }
        }
        runtime.Goexit()
}
