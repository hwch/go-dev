package main

import "fmt"

func max_of_slice(v []int) int {
        if len(v) == 0 {
                return 0
        }
        x := v[0]
        for i := 1; i < len(v); i++ {
                if x < v[i] {
                        x = v[i]
                }
        }
        return x
}

func min_of_slice(v []int) int {
        if len(v) == 0 {
                return 0
        }
        x := v[0]
        for i := 1; i < len(v); i++ {
                if x > v[i] {
                        x = v[i]
                }
        }
        return x
}

func main() {
        x := []int{32, 12, 23, 78, 1}
        fmt.Printf("The %v maximum is %d, minimum is %d\n", x, max_of_slice(x), min_of_slice(x))
}
