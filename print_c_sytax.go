package main

import (
        "bufio"
        "fmt"
        "os"
        "runtime"
        "strings"
        //"time"
        "unicode"
)

func analyse_c_sytax(s string, sk *Stack) error {
        con := -1
        tmp := s
        for con == -1 {
                s, con = get_identifier(tmp, sk)
                tmp = s
        }
        if con == 1 {
                return fmt.Errorf("%s", "未找到标识符")
        }
        if sk.NumOfStack() == 1 {
                return fmt.Errorf("%s", "语法错误")
        }
        id := sk.Pop()
        fmt.Printf("%s是一个", id)
        for sk.Empty() != true {
                fmt.Printf("+++%s+++\n", sk.Pop())
        }
        return nil
}

func get_identifier(s string, sk *Stack) (rs string, ri int) {
        i := 0
        if s == "" {
                ri = 1
                return
        }
        iLen := len(s)
        // fmt.Printf("@@@%s@@@\n", s)

        con := true

        for i < iLen && con {
                switch {
                case unicode.IsLetter(rune(s[i])), unicode.IsNumber(rune(s[i])), s[i] == '_':
                        i++
                        for i < iLen {
                                if unicode.IsLetter(rune(s[i])) || unicode.IsNumber(rune(s[i])) || s[i] == '_' {
                                        i++
                                } else {
                                        break
                                }
                        }
                        con = false
                case s[i] == ' ', s[i] == '\t':
                        i++
                default:
                        con = false
                        i++
                }
        }
        // time.Sleep(time.Second)
        ts := s[:i]
        s1 := strings.Trim(ts, " ")
        s = s[i:]
        sk.Push(s1)
        rs = s
        switch s1 {
        case "":
                ri = -1
        case "char", "short", "int", "long", "float", "double", "auto", "signed", "unsigned", "const", "volatile", "static", "enum", "struct", "union", "void":
                ri = -1
        default:
                if !unicode.IsLetter(rune(s1[0])) && s1[0] != '_' {
                        ri = -1
                } else {
                        ri = 0
                }
        }
        return
        // fmt.Printf("<<<%d|%s|%s|%d>>>\n", iLen, ts[:i], ts[i:], i)

}

type Stack struct {
        cap     int
        data    []string
        pos     int
}

func NewStack(cap int) *Stack {
        var v Stack
        v.data = make([]string, cap)
        v.cap = cap
        v.pos = -1
        return &v
}

func (this *Stack) NumOfStack() int {
        return this.pos + 1
}

func (this *Stack) Push(s string) {
        if !this.Full() {
                this.pos++
                this.data[this.pos] = s
        }
}

func (this *Stack) Pop() string {
        var s string
        if !this.Empty() {
                s = this.data[this.pos]
                this.pos--
        }

        return s
}

func (this *Stack) Empty() bool {
        return this.pos == -1
}

func (this *Stack) Full() bool {
        return this.pos == this.cap-1
}

func (this *Stack) ClearStack() {
        this.pos = -1
}

func (this *Stack) Destory() {
        this.data = nil
        this.cap = 0
        this.pos = -1
        runtime.GC()
}
func main() {
        x := NewStack(1024)
        defer x.Destory()
LOOP:
        for {
                fmt.Fprintf(os.Stderr, "%s", "Please input data: ")
                rin := bufio.NewReader(os.Stdin)
                s, err := rin.ReadString('\n')
                if err != nil {
                        break
                }
                ss := strings.Trim(s, "\n\r")
                // ss := string(s[:len(s)-1])
                switch strings.ToUpper(ss) {
                case "EXIT", "QUIT", "Q":
                        break LOOP
                }
                // x.Push(s)
                if err := analyse_c_sytax(ss, x); err != nil {
                        x.ClearStack()
                        fmt.Printf("Error: %v\n", err)
                }
        }
}
