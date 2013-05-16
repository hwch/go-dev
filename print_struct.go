package main

import "fmt"

type TestStruct struct {
        name    string
        age     int
        sex     bool
}

func main() {
        var a TestStruct

        a.age = 12
        a.name = "Lisi"
        a.sex = true

        fmt.Printf("TestStruct a is %+v\n", a)
}
