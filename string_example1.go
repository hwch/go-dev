package main

import (
        "fmt"
        "strings"
)

func main() {
        if strings.HasPrefix("DEVSTAN.dfafd", "DEVSTAN") {
                fmt.Printf("Yes\n")
        } else {
                fmt.Printf("No\n")
        }
}
