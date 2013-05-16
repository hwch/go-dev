package main

import (
        "fmt"
        "strings"
)

func main() {
        s1 := "123:456"
        s2 := strings.Split(s1, ":")
        for _, v := range s2 {
                fmt.Printf("%s\n", v)
        }
        fmt.Printf("%s\n", strings.Split(s1, ":")[1])
        println("SplitAfter")
        s1 = "fsadf123wtrkl123fsfds123pvmv"
        s2 = strings.SplitAfter(s1, "123")
        for _, v := range s2 {
                fmt.Printf("%s\n", v)
        }
}
