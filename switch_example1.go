package main

import "fmt"

func switch_test(s string) int {
        var sum int
        for i := 0; i < len(s); i = i + 1 {
                fmt.Printf("%c", s[i])
                switch {
                case s[i] >= '0' && s[i] <= '9':
                        sum = sum*10 + int(s[i]-'0')
                case s[i] >= 'a' && s[i] <= 'f':
                        sum = sum*16 + int(s[i]-'a') + 10
                case s[i] >= 'A' && s[i] <= 'F':
                        sum = sum*16 + int(s[i]-'A') + 10
                default:
                        continue
                }
        }
        println(" end")
        return sum
}

func main() {
        fmt.Printf("%d\n", switch_test("123456"))
        fmt.Printf("%d\n", switch_test("aBCdef"))
        fmt.Printf("%d\n", switch_test("1qw23"))
        fmt.Printf("%d\n", switch_test("1aA"))
}
