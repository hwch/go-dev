package main

import (
        "fmt"
)

func MapInt(f func(int) int, arg []int) []int {
        x := make([]int, len(arg))
        for i, n := range arg {
                x[i] = f(n)
        }

        return x
}

func CallbackInt(n int) int {
        return n
}

func MapString(f func(string) string, arg []string) []string {
        x := make([]string, len(arg))

        for i, n := range arg {
                x[i] = f(n)
        }

        return x
}

func CallbackString(s string) string {
        return s
}

func main() {
        /* y := []int { 1, 2, 3, 4, 5 }
           x := MapInt(CallbackInt, y)
           fmt.Printf("%v\n", x) */

        /* y := []int { 1, 2, 3, 4, 5 }
           f := func(n int) int {
           	return n * n
           }
           x := MapInt(f, y)
           fmt.Printf("%v\n", x) */

        /* y := []string { "a", "b", "c", "d", "e" }
           x := MapString(CallbackString, y)
           fmt.Printf("%v\n", x) */

        y := []string{"a", "b", "c", "d", "e"}
        f := func(s string) string {
                return s + ":CallBack"
        }
        x := MapString(f, y)
        fmt.Printf("%v\n", x)
}
