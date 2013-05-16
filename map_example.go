package main

import "fmt"

type mySt struct {
        d       int
        s       string
}

func arg_map(x *map[string]string) {
        (*x)["1"] = "Hello World"
}
func arg_map1(x map[string]string) {
        x["1"] = "Hello World"
}
func main() {
        var x map[string]string

        x = make(map[string]string)

        x["2"] = "Yes"
        x["3"] = "No"
        arg_map(&x)
        fmt.Printf("%s--%s--%s\n", x["2"], x["1"], x["3"])
}
