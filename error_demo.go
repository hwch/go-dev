package main

import (
        "errors"
        "fmt"
)

func TestErrorDemo(a int) error {
        if a < 0 {
                err := fmt.Sprintf("The argument[%v] is not negative number!!!", a)
                return errors.New(err)
        }
        println(a)
        return nil
}

func main() {
        if err := TestErrorDemo(125); err != nil {
                fmt.Printf("%v\n", err)
                return
        }

        if err := TestErrorDemo(-25); err != nil {
                fmt.Printf("%v\n", err)
                return
        }
        fmt.Printf("Hello World\n")
}
