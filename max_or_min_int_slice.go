package main

import "fmt"

func MaxOfIntSlice(val []int) int {
        max := val[0]
        for _, n := range val {
                if max < n {
                        max = n
                }
        }
        return max
}

func MinOfIntSlice(val []int) int {
        min := val[0]
        for _, n := range val {
                if min > n {
                        min = n
                }
        }
        return min
}

func main() {
        x := []int{32, 12, 67, 1, 222, 87}
        fmt.Printf("The maximum and minimum of slice %v is %d and %d\n",
                x, MaxOfIntSlice(x), MinOfIntSlice(x))
}
