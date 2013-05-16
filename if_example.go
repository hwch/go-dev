package main

import "fmt"

func if_test(x bool) int {
        if x == true {
                return 1
        } else {
                return 0
        }
        return 5
}

func main() {
        fmt.Printf("%v\n", if_test(true))
        fmt.Printf("%v\n", if_test(false))
}
