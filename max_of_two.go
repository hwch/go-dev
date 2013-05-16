package main

import "fmt"

func MaxOfTwo(a int, b int) (int, int) {
        if a > b {
                return b, a
        }
        return a, b
}

func main() {
        var a, b, c, d int

        a = 12
        b = 1

        c, d = MaxOfTwo(a, b)

        fmt.Printf("The two number is %d and %d MaxOfTwo return is %d and %d\n",
                a, b, c, d)
}
