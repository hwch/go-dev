package main

import (
        "fmt"
        "strings"
)

func main() {
        if strings.Contains("Hello World", "ello") {
                fmt.Printf("\"ello\" in \"Hello World\"\n")
        }
        if strings.ContainsAny("Hello World", "aqe") {
                fmt.Printf("Yes\n")
        } else {
                fmt.Printf("No\n")
        }
}
