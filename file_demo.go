package main

import (
        "bufio"
        "fmt"
        "os"
)

func main() {
        if len(os.Args) != 2 {
                fmt.Printf("The Arguments is not able to be less than 2\n")
                return
        }

        data := make([]byte, 1024)

        file, err := os.OpenFile(os.Args[1], os.O_RDWR, 0640)
        if err != nil {
                fmt.Printf("%v\n", err)
                return
        }
        defer file.Close()
        rf := bufio.NewReader(file)

        /* wf := bufio.NewWriter(os.Stdout)
           defer wf.Flush() */
        for {
                count, err := rf.Read(data)
                if count == 0 {
                        fmt.Printf("Read complete\n")
                        break
                }
                if err != nil {
                        fmt.Printf("%v\n", err)
                        return
                }
                fmt.Printf("%s", string(data[:count]))
                // wf.Write(data[:count])
                // wf.WriteString(s)
                // fmt.Printf("read:  %s", s)
        }

}
