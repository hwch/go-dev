package main

import (
        "fmt"
        "time"
)

func thread_func(t bool, c chan int) {
        time.Sleep(2 * time.Second)
        if t {
                c <- 2
        } else {
                c <- 1
        }

}

func main() {
        var iTotal int

        c := make(chan int)
        defer close(c)
        iMax := 14
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
        /*
        	for i := 0; i < iMax; i++ {
        		d, ok := <-c
        		if ok == false {
        			break
        		}
        		iTotal += d
        	}
        */
        i := 0
        x := 0
L:
        for {
                d, ok := <-c
                if ok {
                        iTotal += d
                        if i == iMax-1 {
                                break L
                        }
                        i++
                }
                x++
        }

        fmt.Printf("Total is %d, x=%d\n", iTotal, x)
}
