package main

import "fmt"
import "strings"

func main() {
        x := []string{"[a]", "[b]", "[c]", "[d]"}
        y := ""
        for _, v := range x {
                y = y + strings.Trim(v, "[]")
        }
        fmt.Printf("%s\n", y)
}
