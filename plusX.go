package main

import "fmt"

func PlusX(x int) func(int) int {
        return func(y int) int { return y + x }
}

func main() {
        f := PlusX(5)
        fmt.Printf("PlusX(5)= P(9)| %d\n", f(9))
}
