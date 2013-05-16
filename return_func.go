package main

import "fmt"

func plusX(x int) func(int) int {
        return func(y int) int { return x + y }
}

func main() {
        p := plusX(5)

        fmt.Printf("The return function 2 is %d\n", p(2))
}
