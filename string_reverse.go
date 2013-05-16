package main

import "fmt"

func ReverseString(s string) string {
        ret := []byte(s)

        for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
                ret[i], ret[j] = ret[j], ret[i]
        }
        return string(ret)
}

func main() {
        s := "abcdefg"

        fmt.Printf("old string: %s\n"+"new string: %s\n", s, ReverseString(s))
}
