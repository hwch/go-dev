package main

import (
        "fmt"
        "os"
        "path/filepath"
)

func main() {
        s := "*.test"
        s = os.Args[1]
        x, err := filepath.Glob(s)
        if err != nil {
                fmt.Printf("Error:%v\n", err)
                return
        }
        for _, v := range x {
                fmt.Printf("%s\n", v)
        }
}
