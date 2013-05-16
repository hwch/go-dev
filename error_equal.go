package main

import (
        "errors"
        "fmt"
)

func main() {
        var Test_Err = errors.New("Hello World")
        if Test_Err.Error() == "Hello World" {
                fmt.Println("Yes")
        } else {
                fmt.Println("No")
        }
}
