package main

import "fmt"

func PrintArgIntVar(arg ...int) {
        for i, n := range arg {
                fmt.Printf("The %dth argument is %d\n", i, n)
        }
        /* for i := 0; i < len(arg); i++ {
        	fmt.Printf("The %dth argument is %d\n", i, arg[i])
        } */
}

func main() {
        PrintArgIntVar(5, 8, 2, 9, 3, 4, 7, 10)
}
