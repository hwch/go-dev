package main

import "fmt"

func print_int_var(vals ...int) {
        for v, k := range vals {
                fmt.Printf("The %dth is %d\n", v+1, k)
        }
}

func main() {
        print_int_var(9, 8, 7, 6, 5, 4, 3, 2, 1)
}
