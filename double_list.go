package main

import (
        "container/list"
        "fmt"
)

func main() {
        var i int

        v := list.New()

        v.PushBack(1)
        v.PushBack(2)
        v.PushBack(4)

        for x := v.Front(); x != nil; x = x.Next() {
                fmt.Printf("The %dth element is %v\n", i, x.Value)
                i++
        }
}
