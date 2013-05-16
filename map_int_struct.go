package main

import (
        "fmt"
)

type eleMent struct {
        filename string
        fn       func() error
}

func printHello() error {
        fmt.Println("Hello")
        return nil
}
func printWorld() error {
        fmt.Println("World")
        return nil
}
func printYes() error {
        fmt.Println("Yes")
        return nil
}

var testMap = map[int]eleMent{
        1:      eleMent{"printHello", printHello},
        2:      eleMent{"printWorld", printWorld},
        3:      eleMent{"printYes", printYes},
}

func main() {
        for k, v := range testMap {
                fmt.Printf("testMap[%v]=%#v\n", k, v)
        }
        /*
           fmt.Printf("function[%s] output => ", testMap[1].filename)
           testMap[1].fn()
           fmt.Printf("function[%s] output => ", testMap[2].filename)
           testMap[2].fn()
           fmt.Printf("function[%s] output => ", testMap[3].filename)
           testMap[3].fn()
        */
}
