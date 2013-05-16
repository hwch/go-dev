package main

import "fmt"

func main() {
        s := "Hello World"

        fmt.Printf("%s\n", s)

        c := []byte(s)
        c[6] = 0x0
        s = string(c)

        fmt.Printf("%s\n", s)
}
