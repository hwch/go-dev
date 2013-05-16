package main

import (
        "bufio"
        "fmt"
        "iconv"
        "io"
        "os"
)

func main() {
        var fname string

        if len(os.Args) != 2 {
                fname = "test.txt"
        } else {
                fname = os.Args[1]
        }
        f, err := os.Open(fname)
        if err != nil {
                fmt.Printf("Error:%v\n", err)
                return
        }
        defer f.Close()
        rf := bufio.NewReader(f)

        for {
                s, err := rf.ReadString('\n')
                if err != nil && err != io.EOF {
                        fmt.Printf("Error:%v\n", err)
                        return
                }
                if s == "" {
                        break
                }
                sbak := make([]byte, len(s)*2)
                iLen := 0
                if err := iconv.ConvertGBKToUTF8([]byte(s), len(s), sbak, &iLen); err != nil {
                        // if err := iconv.ConvertUTF8ToGBK([]byte(s), len(s), sbak, &iLen); err != nil {
                        fmt.Printf("Error:%v\n", err)
                        return
                }
                fmt.Printf("%s", string(sbak[:iLen]))
        }
}
