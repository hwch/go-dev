package main

import (
        "fmt"

        "strconv"
)

func main() {

here:

        for i := 123; i < 330; i++ {

                x := strconv.Itoa(i) + strconv.Itoa(i*2) + strconv.Itoa(i*3)

                s1 := []byte(x)

                s2 := []byte("123456789")

                for _, v2 := range s2 {

                        isresult := true

                        for _, v1 := range s1 {

                                if v2 == v1 {

                                        isresult = false

                                }

                        }

                        if isresult {

                                continue here

                        }

                }

                fmt.Println(strconv.Itoa(i) + "|" + strconv.Itoa(i*2) + "|" + strconv.Itoa(i*3))

        }

}
