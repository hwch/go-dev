package main

import (
        "errors"
        "fmt"
        "strconv"
        "strings"
)

/* 解析字符串并组为二进制位图 */
func packBit(BitMap []byte, In string) error {
        var (
                i       uint
                j       uint
                pos     uint
        )

        /*
         * b[0]  { 7-0     0, 1, 2, 3, 4, 5, 6, 7 }
         * b[1]  { 15-8    8, 9, 10, 11, 12, 13, 14, 15 }
         * b[2]  { 23-16   16, 17, 18, 19, 20, 21, 22, 23 }
         * b[3]  { 31-24   24, 25, 26, 27, 28, 29, 30 ,31 }
         * b[4]  { 39-32   32, 33, 34, 35, 36, 37, 38, 39 }
         * b[5]  { 47-40   40, 41, 42, 43, 44, 45, 46, 47 }
         * b[6]  { 55-48   48, 49, 50, 51, 52, 53, 54, 55 }
         * b[7]  { 63-56   56, 57, 58, 59, 60, 61, 62, 63 }
         * b[8]  { 71-64   64, 65, 66, 67, 68, 69, 70, 71 }
         * b[9]  { 79-72   72, 73, 74, 75, 76, 77, 78, 79 }
         * b[10] { 87-80   80, 81, 82, 83, 84, 85, 86, 87 }
         * b[11] { 95-88   88, 89, 90, 91, 92, 93, 94, 95 }
         * b[12] { 103-96  96, 97, 98, 99, 100, 101, 102, 103 }
         * b[13] { 111-104 104, 105, 106, 107, 108, 109, 110, 111 }
         * b[14] { 119-112 112, 113, 114, 115, 116, 117, 118, 119 }
         * b[15] { 127-120 120, 121, 122, 123, 124, 125, 126, 127 }
         */
        if BitMap == nil {
                return errors.New("参数不能为空")
        }
        t := strings.Split(In, ",")

        for _, tmp := range t {
                num, err := strconv.Atoi(tmp)
                if err != nil {
                        return err
                }
                pos = uint(num)
                if pos != 0 {
                        pos--
                } else {
                        continue
                }
                i = pos / 8
                j = pos % 8
                BitMap[i] = BitMap[i] | (0x01 << (7 - j))
        }

        return nil
}

const BITMAP_LEN = 17

/* 转换二进制位图到字符数组位图 */
func unPackBit(BitMap []byte) (string, error) {
        //const iMask uint = 0x01
        k := 0
        buf := make([]byte, BITMAP_LEN*8+1)

        if BitMap == nil {
                return "", errors.New("参数不能为空")
        }
        buf[k] = '1'
        k++

        for i := 0; i < BITMAP_LEN-1; i++ {
                for j := 0; j < 8; j++ {
                        if BitMap[i]&(0x01<<uint(7-j)) != 0 {
                                buf[k] = '1'
                        } else {
                                buf[k] = '0'
                        }
                        k++
                }
        }

        myStr := "位图字符串:" + string(buf[:k])
        fmt.Println(myStr)

        return string(buf[:k]), nil
}

func main() {
        x := make([]byte, BITMAP_LEN)
        // err := packBit(x, "0,1,9,17,25,33,41,49,57,65,73,81,89,97,105,113,121,128")
        err := packBit(x, "0,8,16,24,32,40,48,56,64,72,80,88,96,104,112,120,128")
        if err != nil {
                fmt.Printf("Error:%v\n", err)
        } else {
                fmt.Printf("Result:%v\n", x)
        }

        y, err := unPackBit(x)
        if err != nil {
                fmt.Printf("Error:%v\n", err)
        } else {
                fmt.Printf("Result:%v\n", y)
        }
}
