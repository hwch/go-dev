package main

import "fmt"

func example_slice_string(x []string) {

        for i := 0; i < len(x); i++ {
                x[i] = x[i] + " Test"
        }
}

func main() {
        x := []string{"1", "1", "1", "1"}

        example_slice_string(x)

        fmt.Printf("%v\n", x)

}
