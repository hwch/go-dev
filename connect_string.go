package main

import "fmt"

func main() {
        s := "Hello " +
                "World"
        fmt.Printf("%s\n", s)

        s = `Hello
	     World`
        fmt.Printf("%s\n", s)

}
