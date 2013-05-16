package main

import (
        "fmt"
)

func printHello() {
        fmt.Println("Hello")
}
func printWorld() {
        fmt.Println("World")
}
func printYes() {
        fmt.Println("Yes")
}

var testMap = map[string]func(){
        "Hello": printHello,
        "World": printWorld,
        "Yes":   printYes,
}

func main() {
        testMap["Hello"]()
        testMap["World"]()
        testMap["Yes"]()
}
