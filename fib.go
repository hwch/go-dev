package main

import (
        "errors"
        "fmt"
)

/* func Fib(n int) int {
	if n < 2 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
} */

func Fibonacci(n int) ([]int, error) {
        if n < 2 {
                return nil, errors.New("The argument is not less than 2!!!")
        }
        x := make([]int, n)
        x[0], x[1] = 1, 1

        for i := 2; i < n; i++ {
                x[i] = x[i-1] + x[i-2]
        }

        return x, nil
}

func main() {

        // a := Fib(7)
        if a, err := Fibonacci(7); err != nil {
                fmt.Printf("Error:%v\n", err)
                return
        } else {
                fmt.Printf("The fibonacci %d is %v\n", 7, a)
        }
}
