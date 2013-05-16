package main

import "fmt"
import "iconv/better"

func main() {
        x := []byte{0xc4, 0xe3, 0xba, 0xc3, 0xa3, 0xac, 0xca, 0xc0, 0xbd, 0xe7}

        y := make([]byte, 128)
        c, err := better.NewCoder(better.GBK_UTF8_IDX)
        if err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }
        if n, err := c.CodeConvertFunc(x, y); err != nil {
                fmt.Printf("Error: %v\n", err)
        } else {
                fmt.Printf("%s\n", string(y[:n]))
        }
        c1, err1 := better.NewCoder(better.UTF8_UTF16_BE_IDX)
        if err1 != nil {
                fmt.Printf("Error: %v\n", err1)
                return
        }
        y1 := make([]byte, 128)
        x1 := []byte{0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd, 0xef, 0xbc, 0x8c, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c, 0x0a}
        if n, err := c1.CodeConvertFunc(x1, y1); err != nil {
                fmt.Printf("Error: %v\n", err)
        } else {
                for _, v := range y1[:n] {
                        fmt.Printf("0x%02x ", v)
                }
                fmt.Println()
                fmt.Printf("%s\n", string(y1[:n]))
        }
        c2, err2 := better.NewCoder(better.UTF8_UTF16_LE_IDX)
        if err2 != nil {
                fmt.Printf("Error: %v\n", err2)
                return
        }
        y2 := make([]byte, 128)
        x2 := []byte{0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd, 0xef, 0xbc, 0x8c, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c, 0x0a}
        if n, err := c2.CodeConvertFunc(x2, y2); err != nil {
                fmt.Printf("Error: %v\n", err)
        } else {
                for _, v := range y2[:n] {
                        fmt.Printf("0x%02x ", v)
                }
                fmt.Println()
                fmt.Printf("%s\n", string(y2[:n]))
        }
}
