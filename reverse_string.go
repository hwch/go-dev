package main

import "fmt"

func reverse_string(s string) string {
        ret := []byte(s)
        for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
                ret[i], ret[j] = ret[j], ret[i]
        }

        return string(ret)
}

func main() {
        fmt.Printf("abcdefg=>%s\n", reverse_string("abcdefg"))
}
