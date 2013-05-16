package main

import (
        "fmt"
        "strings"
)

func pow_two(v int) int {
        return v * v
}

type hwc interface{}

func map_func_int_example(f func(int) int, v []int) []int {

        x := make([]int, len(v))
        for m, n := range v {
                x[m] = f(n)
        }

        return x
}

func map_func_string_example(f func(string) string, v []string) []string {

        x := make([]string, len(v))
        for m, n := range v {
                x[m] = f(n)
        }

        return x
}

func map_interface_example(f func(hwc) hwc, v []hwc) []hwc {
        x := make([]hwc, len(v))
        for m, n := range v {
                x[m] = f(n)
        }

        return x
}

func multi_type_func(e hwc) hwc {
        switch e.(type) {
        case string:
                return strings.ToUpper(e.(string))
        case int:
                return e.(int) * e.(int)
        default:
                fmt.Printf("Not support this type\n")
        }
        return nil
}

func main() {
        x := []hwc{1, 2, 3, 4, 5, 6, 7, 8, 9}
        y := []hwc{"hello", "world", "you", "are", "welcome"}
        //	f := func(s string) string { return strings.ToUpper(s) }
        //	fmt.Printf("The 7th is %v\n", map_func_int_example(pow_two, x))
        //	fmt.Printf("The 7th is %v\n", map_func_string_example(func(s string) string { return strings.ToUpper(s) }, y))
        fmt.Printf("The 7th is %v\n", map_interface_example(multi_type_func, x))
        fmt.Printf("The 7th is %v\n", map_interface_example(multi_type_func, y))

}
