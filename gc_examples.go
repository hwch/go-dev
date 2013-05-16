package main

import (
        "fmt"
        "runtime"
)

func main() {
        x := make([]byte, 1024)
        for i := 0; i < 1024; i++ {
                x[i] = '0'
        }
        fmt.Println(string(x))
        runtime.GC()
        fmt.Printf("%s", string(x))
}
