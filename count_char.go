package main

import (
        "fmt"
        "unicode/utf8"
)

func main() {
        s := "abcdefg韩"
        s1 := s[:4] + "abc" + s[7:]

        fmt.Printf("The string [%s] has %v unicode characters\n", s, utf8.RuneCount([]byte(s)))
        fmt.Printf("The string [%s] has %v bytes\n", s, len(s))
        fmt.Printf("The old string [%s] new string [%s]\n", s, s1)
}
