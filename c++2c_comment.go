package main

import (
        "bufio"
        "flag"
        "fmt"
        "io"
        "os"
        "path/filepath"
        "runtime"
)

const (
        C_COMMENT_BEGIN = iota
        C_COMMENT_END   = iota
)

func convertCplusplus2C_Comment(s string, isCComment *int) string {
        j := 0
        i := 0
        isCppComment := false

        iLen := len(s)
        buff := make([]byte, iLen+8)

        if *isCComment == C_COMMENT_BEGIN {
                for i < iLen-1 {
                        switch s[i] {
                        case '*':
                                if s[i+1] == '/' {
                                        *isCComment = C_COMMENT_END
                                        buff[j] = s[i]
                                        buff[j+1] = s[i+1]
                                        j += 2
                                        i += 2
                                } else {
                                        buff[j] = s[i]
                                        i++
                                        j++
                                }
                        case '/':
                                switch s[i+1] {
                                case '/':
                                        if *isCComment == C_COMMENT_END {
                                                isCppComment = true
                                                buff[j] = s[i]
                                                buff[j+1] = '*'
                                                i += 2
                                                j += 2

                                                for i < iLen-1 {
                                                        buff[j] = s[i]
                                                        i++
                                                        j++
                                                }
                                        } else {
                                                buff[j] = s[i]
                                                i++
                                                j++
                                        }
                                default:
                                        buff[j] = s[i]
                                        i++
                                        j++
                                }
                        default:
                                buff[j] = s[i]
                                j++
                                i++
                        }

                }
        } else {
                for i < iLen-1 {
                        switch s[i] {
                        case '/':
                                switch s[i+1] {
                                case '*':
                                        *isCComment = C_COMMENT_BEGIN
                                        buff[j] = s[i]
                                        buff[j+1] = s[i+1]
                                        j += 2
                                        i += 2
                                case '/':
                                        if *isCComment == C_COMMENT_END {
                                                isCppComment = true
                                                buff[j] = s[i]
                                                buff[j+1] = '*'
                                                i += 2
                                                j += 2

                                                for i < iLen-1 {
                                                        buff[j] = s[i]
                                                        i++
                                                        j++
                                                }
                                        } else {
                                                buff[j] = s[i]
                                                i++
                                                j++
                                        }
                                default:
                                        buff[j] = s[i]
                                        j++
                                        i++
                                }
                        case '*':
                                if s[i+1] == '/' {
                                        *isCComment = C_COMMENT_END
                                        buff[j] = s[i]
                                        buff[j+1] = s[i+1]
                                        j += 2
                                        i += 2
                                } else {
                                        buff[j] = s[i]
                                        j++
                                        i++
                                }
                        default:
                                buff[j] = s[i]
                                j++
                                i++
                        }
                }
        }
        if isCppComment {
                if s[i-1] == '\r' {
                        buff[j-1] = ' '
                        buff[j] = '*'
                        buff[j+1] = '/'
                        buff[j+2] = '\r'
                        buff[j+3] = '\n'
                        j += 4
                } else {
                        buff[j] = ' '
                        buff[j+1] = '*'
                        buff[j+2] = '/'
                        buff[j+3] = '\n'
                        j += 4
                }
        } else {
                if i < iLen {
                        buff[j] = s[i]
                        j++
                }
        }
        return string(buff[:j])
}

func walkFunc(path string, info os.FileInfo, err1 error) error {
        /* fmt.Printf("FileName[%s]FileSize[%d]Dir[%v]\n",
           filepath.Join(path, info.Name()), info.Size(), info.IsDir()) */

        if info.IsDir() {
                return err1
        }
        // FileName := filepath.Join(path, info.Name())
        FileName := path
        if filepath.Ext(FileName) == ".c" || filepath.Ext(FileName) == ".pc" {
                // fmt.Printf("%s is a C file\n", FileName)
                rf, err := os.Open(FileName)
                if err != nil {
                        return err
                }

                defer rf.Close()
                config_bak := FileName + ".bak_hwch.20130107"
                wf, err := os.OpenFile(config_bak, os.O_CREATE|os.O_RDWR|os.O_TRUNC|os.O_EXCL, 0644)
                if err != nil {
                        if os.IsExist(err) {
                                wf, err = os.OpenFile(config_bak, os.O_TRUNC|os.O_RDWR, 0644)
                                if err != nil {
                                        return err
                                }
                        } else {
                                return err
                        }
                }
                defer wf.Close()
                //      data := make([]byte, 1024)
                rfb := bufio.NewReader(rf)
                wfb := bufio.NewWriter(wf)
                isCComment := C_COMMENT_END

                for {
                        s, err := rfb.ReadString('\n')
                        if err != nil {
                                if err == io.EOF {
                                        if s != "" {
                                                sBak := convertCplusplus2C_Comment(s, &isCComment)
                                                if _, err := wfb.WriteString(sBak); err != nil {
                                                        return err
                                                }
                                        }
                                        break
                                } else {
                                        return err
                                }
                        }
                        sBak := convertCplusplus2C_Comment(s, &isCComment)
                        if _, err := wfb.WriteString(sBak); err != nil {
                                return err
                        }
                }
                wfb.Flush()
                if err := wf.Close(); err != nil {
                        return err
                }
                if err := rf.Close(); err != nil {
                        return err
                }
                if err := os.Remove(FileName); err != nil {
                        return err
                }
                if err := os.Rename(config_bak, FileName); err != nil {
                        return err
                }
                runtime.GC()
        }

        return err1
}

func main() {
        sDir := ""
        flag.StringVar(&sDir, "dir", "", "要查询的目录")
        flag.Parse()
        if sDir == "" {
                flag.PrintDefaults()
                return
        }
        if err := filepath.Walk(sDir, walkFunc); err != nil {
                fmt.Printf("Error: %v\n", err)
        } else {
                fmt.Println("Complete!!!!!!!!!")
        }
}
