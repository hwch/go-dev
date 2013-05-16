package main

import (
        "fmt"
        "strings"
)

func main() {
        s := "dfjklsdajfoiwejciofdHellofjoiajklewoisdmoifjio"

        x := strings.Index(s, "Hello")
        if x < 0 {
                println("Not Found")
                return
        }
        fmt.Printf("%s\n", s[x:])

}
