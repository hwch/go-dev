package main

import (
        "bufio"
        "flag"
        "fmt"
        "io"
        "os"
        "strings"
)

// 从字符串中删除某个字符
func delSomeCh(s string, ch byte) string {
        i := 0
        j := 0

        Len := len(s)
        v := make([]byte, Len)
        for i < Len {
                if s[i] != ch {
                        v[j] = s[i]
                        j++
                }
                i++
        }
        return string(v[:j])
}

func convertLowerFile() {
        rf, err := os.Open(os.Args[1])
        if err != nil {
                fmt.Printf("Error:%v", err)
                return
        }
        defer rf.Close()
        wf, err := os.OpenFile(os.Args[1]+".bak", os.O_CREATE|os.O_RDWR|os.O_TRUNC|os.O_EXCL, 0644)
        if err != nil {
                if os.IsExist(err) {
                        wf, err = os.OpenFile(os.Args[1]+".bak", os.O_TRUNC|os.O_RDWR, 0644)
                        if err != nil {
                                fmt.Printf("Error:%v", err)
                                return
                        }
                } else {
                        fmt.Printf("Error:%v", err)
                        return
                }
        }
        defer wf.Close()

        rfb := bufio.NewReader(rf)
        wfb := bufio.NewWriter(wf)
        defer wfb.Flush()
        for {
                s, err := rfb.ReadString('\n')
                if err != nil {
                        if err != io.EOF {
                                fmt.Printf("Error:%v", err)
                                return
                        }
                        break
                }
                switch {
                case s[0] == '\n':
                        fallthrough
                case strings.Contains(s, "g_Gbk2Unicode"):
                        fallthrough
                case strings.Contains(s, "g_Unicode2Gbk"):
                        fallthrough
                case strings.HasPrefix(s, "//"):
                        if _, err := wfb.WriteString(s); err != nil {
                                fmt.Printf("Error:%v\n", err)
                                return
                        }
                default:
                        if _, err := wfb.WriteString(strings.ToLower(s)); err != nil {
                                fmt.Printf("Error:%v\n", err)
                                return
                        }
                }
        }
}

func convertStandardFile() {
        rf, err := os.Open(os.Args[1])
        if err != nil {
                fmt.Printf("Error:%v", err)
                return
        }
        defer rf.Close()
        wf, err := os.OpenFile("."+os.Args[1]+".bak", os.O_CREATE|os.O_RDWR|os.O_TRUNC|os.O_EXCL, 0644)
        if err != nil {
                if os.IsExist(err) {
                        wf, err = os.OpenFile("."+os.Args[1]+".bak", os.O_TRUNC|os.O_RDWR, 0644)
                        if err != nil {
                                fmt.Printf("Error:%v", err)
                                return
                        }
                } else {
                        fmt.Printf("Error:%v", err)
                        return
                }
        }
        defer wf.Close()

        rfb := bufio.NewReader(rf)
        wfb := bufio.NewWriter(wf)
        defer wfb.Flush()
        BLOOP := true
        for BLOOP {
                strBak := "        "
                for ih := 0; ih < 4; ih++ {
                        s, err := rfb.ReadString('\n')
                        if err != nil && err != io.EOF {
                                fmt.Printf("Error:%v", err)
                                return
                        }
                        if s == "" {
                                BLOOP = false
                                break
                        }
                        if s[0] == '\n' {
                                continue
                        }
                        tmpStr := delSomeCh(s, '\n')
                        strBak = strBak + tmpStr + ", "

                }
                wfb.WriteString(strBak + "\n")
        }
}

func U2G() {
        rf, err := os.Open(os.Args[1])
        if err != nil {
                fmt.Printf("Error:%v", err)
                return
        }
        defer rf.Close()
        wf, err := os.OpenFile("unicode2gbk.value", os.O_CREATE|os.O_RDWR|os.O_TRUNC|os.O_EXCL, 0644)
        if err != nil {
                if os.IsExist(err) {
                        wf, err = os.OpenFile("unicode2gbk.value", os.O_TRUNC|os.O_RDWR, 0644)
                        if err != nil {
                                fmt.Printf("Error:%v", err)
                                return
                        }
                } else {
                        fmt.Printf("Error:%v", err)
                        return
                }
        }
        defer wf.Close()

        rfb := bufio.NewReader(rf)
        wfb := bufio.NewWriter(wf)
        defer wfb.Flush()
        if _, err := wfb.WriteString("var g_Unicode2Gbk = map[uint16]uint16{\n"); err != nil {
                fmt.Printf("Error:%v", err)
                return
        }
        BLOOP := true
        for BLOOP {
                strBak := "        "
                for ih := 0; ih < 4; ih++ {
                        s, err := rfb.ReadString('\n')
                        if err != nil && err != io.EOF {
                                fmt.Printf("Error:%v", err)
                                return
                        }
                        if s == "" {
                                BLOOP = false
                                break
                        }
                        if s[0] == '\n' {
                                continue
                        }
                        tmpStr := delSomeCh(s, '\n')
                        vStr := strings.Split(tmpStr, "\t")
                        strBak = strBak + fmt.Sprintf("0x%s:0x%s, ", strings.ToLower(vStr[1]), strings.ToLower(vStr[0]))

                }
                wfb.WriteString(strBak + "\n")
        }

        if _, err := wfb.WriteString("}\n"); err != nil {
                fmt.Printf("Error:%v", err)
                return
        }
}
func G2U() {
        rf, err := os.Open(os.Args[1])
        if err != nil {
                fmt.Printf("Error:%v", err)
                return
        }
        defer rf.Close()
        wf, err := os.OpenFile("gbk2unicode.value", os.O_CREATE|os.O_RDWR|os.O_TRUNC|os.O_EXCL, 0644)
        if err != nil {
                if os.IsExist(err) {
                        wf, err = os.OpenFile("gbk2unicode.value", os.O_TRUNC|os.O_RDWR, 0644)
                        if err != nil {
                                fmt.Printf("Error:%v", err)
                                return
                        }
                } else {
                        fmt.Printf("Error:%v", err)
                        return
                }
        }
        defer wf.Close()

        rfb := bufio.NewReader(rf)
        wfb := bufio.NewWriter(wf)
        defer wfb.Flush()
        if _, err := wfb.WriteString("var g_Gbk2Unicode = map[uint16]uint16{\n"); err != nil {
                fmt.Printf("Error:%v", err)
                return
        }
        BLOOP := true
        for BLOOP {
                strBak := "        "
                for ih := 0; ih < 4; ih++ {
                        s, err := rfb.ReadString('\n')
                        if err != nil && err != io.EOF {
                                fmt.Printf("Error:%v", err)
                                return
                        }
                        if s == "" {
                                BLOOP = false
                                break
                        }
                        if s[0] == '\n' {
                                continue
                        }
                        tmpStr := delSomeCh(s, '\n')
                        vStr := strings.Split(tmpStr, "\t")
                        strBak = strBak + fmt.Sprintf("0x%s:0x%s, ", strings.ToLower(vStr[0]), strings.ToLower(vStr[1]))

                }
                wfb.WriteString(strBak + "\n")
        }
        if _, err := wfb.WriteString("}\n"); err != nil {
                fmt.Printf("Error:%v", err)
                return
        }
}
func main() {
        flag.Parse()
        if flag.NArg() != 1 {
                fmt.Printf("参数个数不对\n")
                return
        }
        // convertStandardFile()
        // convertLowerFile()
        // convertStandardFile()
        U2G()
        G2U()
}
