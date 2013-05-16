package main

import "fmt"

func fibbonacci(v int) []int {
        if v < 2 {
                return nil
        }
        x := make([]int, v)
        x[0], x[1] = 1, 1

        for i := 2; i < v; i++ {
                x[i] = x[i-1] + x[i-2]
        }

        return x
}

func main() {
        fmt.Printf("The 2th fib is %v\n", fibbonacci(2))
        fmt.Printf("The 5th fib is %v\n", fibbonacci(5))
        fmt.Printf("The 7th fib is %v\n", fibbonacci(7))
        fmt.Printf("The 12th fib is %v\n", fibbonacci(12))
}
