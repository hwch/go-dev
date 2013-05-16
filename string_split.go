package main

import (
        "fmt"
        "strings"
)

func main() {
        x := "1,2,3,4,5"
        for _, y := range strings.SplitN(x, ",", 2) {
                fmt.Printf("%s\n", y)
        }
}
