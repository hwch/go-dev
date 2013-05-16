package main

import (
        "fmt"
        "istack"
)

/* import (
	"fmt"
	"errors"
	"strconv"
)

type stack struct {
	pos     int
	data    [10]int
}

func (s *stack) StackInit() {
	//func (s stack) StackInit() {
	s.pos = -1
	for i := 0; i < len(s.data); i++ {
		s.data[i] = 0
	}
}

func (s *stack) Pop() (int, error) {
	//func (s stack) Pop() (int, error) {

	if s.pos == -1 {
		err := fmt.Sprintf("The stack is empty")
		return s.pos, errors.New(err)
	}
	s.pos--
	return s.data[s.pos+1], nil
}

func (s *stack) Push(val int) error {
	//func (s stack) Push(val int) error {
	if s.pos == len(s.data)-1 {
		err := fmt.Sprintf("The stack is full")
		return errors.New(err)
	}
	s.pos++
	s.data[s.pos] = val
	return nil
}
func (s *stack) String() string {
	var i int
	str := ""

	for i = 0; i < len(s.data)-1; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "] " // fmt.Sprintf("[%v:%v] ", i, s.data[i])
	}

	str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]" // fmt.Sprintf("[%v:%v] ", i, s.data[i])
	return str
} */

func main() {
        if s, err := istack.New(8); /*New(8)*/ err != nil {
                fmt.Printf("Stack initialize failed !!!\n")
                return
        } else {
                //var s stack
                //        for i := 0; i < 8; i++ {
                for i := 1024; i > 0; i-- {
                        if bl := s.Push(i); bl != nil {
                                fmt.Printf("Error:%v\n", bl)
                                break
                        } else {
                                fmt.Printf("The in value is %v\n", i)
                        }
                }
                fmt.Printf("%v\n", s.String())
                //	fmt.Printf("%v\n", s)
                //        fmt.Printf("%v\n", s.pos)
                //        for i := 0; i < 8; i++ {
                for {
                        if val, bl := s.Pop(); bl != nil {
                                fmt.Printf("Error:%v\n", bl)
                                break
                        } else {
                                fmt.Printf("The out value is %v\n", val)
                        }
                }
        }

}
