package main

import (
        "bytes"
        "encoding/binary"
        "flag"
        "fmt"
        "io"
        "os"
        "runtime"
        "strconv"
        "unsafe"
)

var fileName string

const (
        BUFFER_READ_MAX_BYTES = 8192
)

func init() {
        flag.StringVar(&fileName, "f", "", "The file name of archive")
        flag.Parse()
}

func isLittleEndian() bool {
        var x int32 = 0x12345678
        p := unsafe.Pointer(&x)
        p1 := (*[4]byte)(p)
        return p1[0] == 0x78
}

func main() {
        var memNum uint32
        var longFileName []byte
        iLen := 0
        if fileName == "" {
                flag.Usage()
                return
        }

        rf, err := os.Open(fileName)
        if err != nil {
                _, _, LineNo, _ := runtime.Caller(0)
                fmt.Printf("Line[%d], Error: %v\n", LineNo, err)
                return
        }
        defer rf.Close()
        aheader := make([]byte, 8)
        if _, err := rf.Read(aheader); err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }
        mheader := make([]byte, 60)
        if n, err := rf.Read(mheader); err != nil {
                _, _, LineNo, _ := runtime.Caller(0)
                fmt.Printf("Line[%d], Error: %v\n", LineNo, err)
                return
        } else {
                if n != 60 {
                        _, _, LineNo, _ := runtime.Caller(0)
                        fmt.Printf("Line[%d], Error: File format is incorrect\n", LineNo)
                        return
                }
                iLen := 0
                if v, err := strconv.Atoi(string(bytes.Trim(mheader[48:58], " "))); err != nil {
                        fmt.Printf("Error: %v\n", err)
                        return
                } else {
                        iLen = v
                        if iLen%2 == 1 {
                                iLen++
                        }
                }
                mcon := make([]byte, iLen)
                if n1, err := rf.Read(mcon); err != nil {
                        _, _, LineNo, _ := runtime.Caller(0)
                        fmt.Printf("Line[%d], Error: %v\n", LineNo, err)
                        return
                } else {
                        if n1 != iLen {
                                _, _, LineNo, _ := runtime.Caller(0)
                                fmt.Printf("Line[%d], Error: File format is incorrect\n", LineNo)
                                return
                        }
                }

                myStr := "所有符号为: "
                memNum = binary.BigEndian.Uint32(mcon)
                i := 4 + 4*int(memNum)
                x := bytes.Split(mcon[i:], []byte{0})
                for _, v := range x {
                        if len(v) == 0 {
                                continue
                        }
                        myStr = myStr + fmt.Sprintf("[%s]", string(v))
                }
                fmt.Println(myStr)
        }

        for {
                if n, err := rf.Read(mheader); err != nil {
                        if err == io.EOF {
                                break
                        }
                        _, _, LineNo, _ := runtime.Caller(0)
                        fmt.Printf("Line[%d], Error: %v\n", LineNo, err)
                        return
                } else {
                        if n != 60 {
                                _, _, LineNo, _ := runtime.Caller(0)
                                fmt.Printf("Line[%d], Error: File format is incorrect\n", LineNo)
                                return
                        }

                        fmt.Print("Read 60 bytes:")
                        myStr := ""
                        for _, v := range mheader {
                                if strconv.IsPrint(rune(v)) {
                                        // myStr = myStr + fmt.Sprintf("%02x", v)
                                        myStr = myStr + fmt.Sprintf("%c", v)
                                } else {
                                        myStr = myStr + fmt.Sprintf("<%d>", v)
                                }

                        }
                        fmt.Println(myStr)
                        if v, err := strconv.Atoi(string(bytes.Trim(mheader[48:58], " "))); err != nil {
                                _, _, LineNo, _ := runtime.Caller(0)
                                fmt.Printf("Line[%d], Error: %v\n", LineNo, err)
                                return
                        } else {
                                iLen = v
                                if iLen%2 == 1 {
                                        iLen++
                                }
                        }
                }
                if string(mheader[:2]) == "//" {
                        longFileName = make([]byte, iLen)
                        if _, err := rf.Read(longFileName); err != nil {
                                _, _, LineNo, _ := runtime.Caller(0)
                                fmt.Printf("Line[%d], Error: %v\n", LineNo, err)
                                return
                        }
                        continue
                }
                counter := 0
                mcon := make([]byte, BUFFER_READ_MAX_BYTES)
                if iLen < BUFFER_READ_MAX_BYTES {
                        mcon := make([]byte, iLen)
                        if n, err := rf.Read(mcon); err != nil {
                                _, _, LineNo, _ := runtime.Caller(0)
                                fmt.Printf("Line[%d], Error: %v\n", LineNo, err)
                                return
                        } else {
                                if n != iLen {
                                        _, _, LineNo, _ := runtime.Caller(0)
                                        fmt.Printf("Line[%d], Error: File format is incorrect\n", LineNo)
                                        return
                                }
                        }
                } else {
                        for iLen > counter {
                                if n, err := rf.Read(mcon); err != nil {
                                        _, _, LineNo, _ := runtime.Caller(0)
                                        fmt.Printf("Line[%d], Error: %v\n", LineNo, err)
                                        return
                                } else {
                                        counter += n
                                        if iLen-counter < BUFFER_READ_MAX_BYTES {
                                                mcon1 := make([]byte, iLen-counter)
                                                if n, err := rf.Read(mcon1); err != nil {
                                                        _, _, LineNo, _ := runtime.Caller(0)
                                                        fmt.Printf("Line[%d], Error: %v\n", LineNo, err)
                                                        return
                                                } else {
                                                        if n != iLen-counter {
                                                                _, _, LineNo, _ := runtime.Caller(0)
                                                                fmt.Printf("Line[%d], Error: File format is incorrect\n", LineNo)
                                                                return
                                                        }
                                                        break
                                                }
                                        }
                                }

                        }
                }

                runtime.GC()
                iLen = 0
        }
        myStr := ""
        for _, v := range longFileName {
                if strconv.IsPrint(rune(v)) {
                        // myStr = myStr + fmt.Sprintf("%02x", v)
                        myStr = myStr + fmt.Sprintf("%c", v)
                } else {
                        myStr = myStr + fmt.Sprintf("<%d>", v)
                }

        }
        fmt.Printf("longFileName: %s\n", myStr)
}
