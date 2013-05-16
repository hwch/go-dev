package main

import "fmt"

func Bubble(v []int) []int {
        if v == nil {
                return nil
        }
        for {
                swapped := false
                for i := 1; i < len(v); i++ {
                        if v[i-1] > v[i] {
                                v[i-1], v[i] = v[i], v[i-1]
                                swapped = true
                        }
                }
                if swapped == false {
                        break
                }
        }

        return v
}

func main() {
        x := []int{21, 3, 98, 1, 100, 30, 11}

        fmt.Printf("Before:%v, After:%v\n", x, Bubble(x))
}
