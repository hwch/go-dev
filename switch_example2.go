package main

import "fmt"

func switch_example1(x int) {
        fmt.Printf("switch_example1: ")
        switch x {
        case 0:
        case 1:
                fmt.Printf("0-1\n")
        default:
                fmt.Printf("2-9\n")
        }
}

func switch_example2(x int) {
        fmt.Printf("switch_example2: ")
        switch x {
        case 0:
                fallthrough
        case 1:
                fmt.Printf("0-1\n")
        default:
                fmt.Printf("2-9\n")
        }
}

func main() {

        switch_example1(0)
        switch_example2(0)
}
