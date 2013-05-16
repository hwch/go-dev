package main

import (
        "errors"
        "fmt"
        "strconv"
)

const MAX_ELEMENTS = 1024

type HStack struct {
        pos     int
        data    [MAX_ELEMENTS]int
        max_num int
}

func NewHStacker(max int) (*HStack, error) {

        if max <= 0 || max > MAX_ELEMENTS {
                return nil, errors.New("Argument is too large !")
        }
        x := new(HStack)
        x.pos = -1
        x.max_num = max
        return x, nil
}

func (hs *HStack) IsEmpty() bool {
        if hs.pos == -1 {
                return true
        }
        return false
}

func (hs *HStack) IsFull() bool {
        if hs.pos == hs.max_num-1 {
                return true
        }
        return false
}

func (hs *HStack) Pop() int {
        v := hs.pos
        hs.pos--
        return hs.data[v]
}

func (hs *HStack) Push(v int) {
        hs.pos++
        hs.data[hs.pos] = v
}

func (hs HStack) String() string {
        str := ""
        for i := 0; i < hs.pos; i++ {
                str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(hs.data[i]) + "]"
        }
        return str
}

func main() {
        v, err := NewHStacker(8)
        //	v, err := NewHStacker(1025)
        if err != nil {
                fmt.Printf("Initialize stack failed: %v\n", err)
                return
        }
        fmt.Printf("In => ")
        for i := 0; i < 1024; i++ {
                if v.IsFull() {
                        break
                }
                if i == 5 {
                        fmt.Printf("   String of stack <==> %v\n", v)
                }
                v.Push(i * i)
                fmt.Printf("[ %d ]", i*i)
        }
        fmt.Printf("\n")
        fmt.Printf("String of stack <==> %v\n", v)
        fmt.Printf("Out => ")
        for {
                if v.IsEmpty() {
                        break
                }
                fmt.Printf("[ %d ]", v.Pop())
        }
        fmt.Printf("\n")
}
