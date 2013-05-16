package main

import (
        "fmt"
        "strconv"
)

func main() {
        s1 := "10111101"

        for x, y := range s1 {
                myStr := strconv.Itoa(x) + "==>"
                if y == '0' {
                        myStr = myStr + "无"
                } else {
                        myStr = myStr + "有"
                }
                fmt.Printf("%s\n", myStr)
        }
}
