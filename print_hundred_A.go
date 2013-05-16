package main

import "fmt"

func print_A() {
        s := "A"
        for i := 0; i < 100; i = i + 1 {
                fmt.Printf("%s\n", s)
                s = s + "A"
        }
}

func main() {
        print_A()
}
