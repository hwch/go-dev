package main

import (
        "fmt"
        "unicode/utf8"
)

func count_string_byte(s string) (int, int) { // x is the number of characters in s, y is the byte of s
        var x, y int

        x = len([]byte(s))
        y = utf8.RuneCountInString(s)
        // y = utf8.RuneCount([]byte(s))

        return x, y
}

func main() {
        x, y := count_string_byte("1234qaz")
        fmt.Printf("number[%v], byte[%v]\n", x, y)
}
