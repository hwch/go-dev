package main

import (
        "fmt"
        "time"
)

func main() {
        fmt.Printf("Time now %s\n", time.Now().Format(time.ANSIC))
}
