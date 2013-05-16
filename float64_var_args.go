package main

import "fmt"

func float64_slice_average(f ...float64) float64 {
        var ret float64
        var i int

        for _, v := range f {
                ret += v
                i++
        }

        if i == 0 {
                return 0.0
        }
        return ret / float64(i)
}

func main() {
        // f := [...]float64{ 1.2, 3.2, 5.6, 7.8, 2.4, 4.4 }
        // fmt.Printf("%v\n", float64_slice_average(f[:]))
        fmt.Printf("%v\n", float64_slice_average(1.2, 3.2, 5.6, 7.8, 2.4, 4.4))
}
