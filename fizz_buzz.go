package main

import "fmt"

func fizz_buzz() {
        for i := 1; i <= 100; i = i + 1 {
                switch {
                case i%15 == 0:
                        fmt.Printf(" FizzBuzz[%d]\n", i)
                case i%3 == 0:
                        fmt.Printf(" Fizz[%d] ", i)
                case i%5 == 0:
                        fmt.Printf(" Buzz[%d] ", i)
                }
        }
}

func main() {
        fizz_buzz()
        fmt.Println()
}
