package main

import "fmt"

func ret_correct_seq(a, b int) (int, int) {
        if a < b {
                println("a < b")
                return a, b
        }
        println("a >= b")
        return b, a
}

/* func ret_correct_seq(a, b int) (c, d int) {
	if a < b {
		println("a < b")
		c, d = a, b
	} else {
		println("a >= b")
		c, d = b, a
	}
	return
} */

func main() {
        x, y := 1, 3

        fmt.Printf("Before %v, %v\n", x, y)
        x, y = ret_correct_seq(x, y)
        fmt.Printf("After %v, %v\n", x, y)
}
