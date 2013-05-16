package main

import (
        "bufio"
        "fmt"
        "os"
        "os/signal"
)

func main() {
        c := make(chan os.Signal)
        signal.Notify(c)
        go func(sig chan os.Signal) {
                for {
                        c := <-sig
                        if os.Interrupt != c {
                                fmt.Printf("Not CTRL+C \n")
                                os.Exit(0)
                        }       /* else {
                                fmt.Printf("CTRL+C \n")
                        }*/
                }
        }(c)
        stdin := bufio.NewReader(os.Stdin)
        for {
                x, _ := stdin.ReadSlice('\n')
                fmt.Printf("%v\n", x)
        }
}
