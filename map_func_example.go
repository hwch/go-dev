package main

import (
        "fmt"
        "strings"
)

func pow_two(v int) int {
        return v * v
}

func map_func_int_example(f func(int) int, v []int) []int {

        x := make([]int, len(v))
        for m, n := range v {
                x[m] = f(n)
        }

        return x
}

func map_func_string_example(f func(string) string, v []string) []string {

        x := make([]string, len(v))
        for m, n := range v {
                x[m] = f(n)
        }

        return x
}

func main() {
        x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
        y := []string{"hello", "world", "you", "are", "welcome"}
        //	f := func(s string) string { return strings.ToUpper(s) }
        fmt.Printf("The 7th is %v\n", map_func_int_example(pow_two, x))
        fmt.Printf("The 7th is %v\n", map_func_string_example(func(s string) string { return strings.ToUpper(s) }, y))
}
