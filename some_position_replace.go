package main

import "fmt"

func some_position_replace(src, dest string, pos int) string {
        var i, j int
        z := []byte(src)
        if pos > len(z) || pos <= 0 {
                return src
        }
        for i < pos-1 {
                i = i + 1
        }
        for j = 0; j < len(dest); j = j + 1 {
                if i >= len(z) {
                        break
                }
                z[i] = dest[j]
                i = i + 1
        }

        return string(z)
}

func main() {
        fmt.Printf("%s\n", some_position_replace("1234567890", "abc", 1))
}
