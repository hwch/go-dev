package main

import "fmt"
import "time"

func main() {
        for {
                fmt.Println("Hello World")
                time.Sleep(10 * time.Second)
        }
}
