package main

import (
        "fmt"
        "time"
)

func main() {
        vNow := time.Now()
        vTime := fmt.Sprintf("%02d%02d%02d", vNow.Hour(), vNow.Minute(), vNow.Second())
        vDate := fmt.Sprintf("%04d%02d%02d", vNow.Year(), vNow.Month(), vNow.Day())
        fmt.Printf("%s%s\n", vDate, vTime)
}
