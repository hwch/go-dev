package main

import "fmt"

func main() {
        i := 0

Here:
        fmt.Printf("This is %d\n", i)
        i = i + 1
        if i < 10 {
                goto Here
        }

}
