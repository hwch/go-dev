package main

import "fmt"
import "errors"

func byteCopy(dest, src []byte, iLen int) {
        for i := 0; i < iLen; i++ {
                dest[i] = src[i]
        }

}

func getKeyAndValue(s string, k, v *string) error {
        var i int
        var j int

        iLen := len(s)
        vk := make([]byte, iLen)
        vv := make([]byte, iLen)
        *k = ""
        *v = ""
        if s[0] == '=' {
                return errors.New("缺少左值")
        }
        for i < iLen {
                if s[i] == '=' {
                        break
                }
                vk[j] = s[i]
                i++
                j++
        }
        *k = string(vk[:j])
        i++
        if i >= iLen {
                return errors.New("缺少右值")
        }
        j = 0
        for i < iLen {
                vv[j] = s[i]
                i++
                j++
        }
        *v = string(vv[:j])

        return nil
}

func main() {
        s := "1"
        s1 := "2=3"
        s2 := "4="

        k := ""
        v := ""

        getKeyAndValue(s, &k, &v)
        fmt.Printf("%s|%s\n", k, v)
        getKeyAndValue(s1, &k, &v)
        fmt.Printf("%s|%s\n", k, v)
        getKeyAndValue(s2, &k, &v)
        fmt.Printf("%s|%s\n", k, v)
}
