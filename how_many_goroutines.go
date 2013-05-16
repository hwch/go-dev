package main

import "fmt"

func thread_func(t bool, c chan int) {
        if t {
                c <- 1
        } else {
                c <- 0
        }
}

func main() {
        var iTotal int

        c := make(chan int)
        defer close(c)
        iMax := 11
        for i := 0; i < iMax; i++ {
                if i%2 == 0 {
                        go thread_func(false, c)
                } else {
                        go thread_func(true, c)
                }
        }
        /*
        	i := 0
        	if iMax == 0 {
        		return
        	}
        	L: for {
        		select {
        		case d := <-c:
        			iTotal += d
        			if i == iMax-1 {
        				break L
        			}
        			i++
        		}
        	}
        */

        for i := 0; i < iMax; i++ {
                iTotal += <-c
        }

        fmt.Printf("Total is %d\n", iTotal)
}
