package main

import (
        "bytes"
        "fmt"
        "iconv"
        "os"
        "path"
)

var LPInfo = map[byte]string{
        0:      "Blues", 1: "ClassicRock", 2: "Country", 3: "Dance",
        4:      "Disco", 5: "Funk", 6: "Grunge", 7: "Hip-Hop",
        8:      "Jazz", 9: "Metal", 10: "NewAge", 11: "Oldies",
        12:     "Other", 13: "Pop", 14: "R&B", 15: "Rap",
        16:     "Reggae", 17: "Rock", 18: "Techno", 19: "Industrial",
        20:     "Alternative", 21: "Ska", 22: "DeathMetal", 23: "Pranks",
        24:     "Soundtrack", 25: "Euro-Techno", 26: "Ambient", 27: "Trip-Hop",
        28:     "Vocal", 29: "Jazz+Funk", 30: "Fusion", 31: "Trance",
        32:     "Classical", 33: "Instrumental", 34: "Acid", 35: "House",
        36:     "Game", 37: "SoundClip", 38: "Gospel", 39: "Noise",
        40:     "AlternRock", 41: "Bass", 42: "Soul", 43: "Punk",
        44:     "Space", 45: "Meditative", 46: "InstrumentalPop", 47: "InstrumentalRock",
        48:     "Ethnic", 49: "Gothic", 50: "Darkwave", 51: "Techno-Industrial",
        52:     "Electronic", 53: "Pop-Folk", 54: "Eurodance", 55: "Dream",
        56:     "SouthernRock", 57: "Comedy", 58: "Cult", 59: "Gangsta",
        60:     "Top40", 61: "ChristianRap", 62: "Pop/Funk", 63: "Jungle",
        64:     "NativeAmerican", 65: "Cabaret", 66: "NewWave", 67: "Psychadelic",
        68:     "Rave", 69: "Showtunes", 70: "Trailer", 71: "Lo-Fi",
        72:     "Tribal", 73: "AcidPunk", 74: "AcidJazz", 75: "Polka",
        76:     "Retro", 77: "Musical", 78: "Rock&Roll", 79: "HardRock",
}

// 判断IDv3 Version2还是Version1
func isAtHeadOfMP3(fStr []byte) bool {
        if bytes.Equal(fStr[:3], []byte("ID3")) {
                return true
        }
        return false
}

// 字节拷贝
func byteCopy(dest, src []byte, iLen int) {
        for i := 0; i < iLen; i++ {
                dest[i] = src[i]
        }

}

// 设置字节
func byteSet(src []byte, s byte, iLen int) {
        for i := 0; i < iLen; i++ {
                src[i] = s
        }
}

func main() {
        var rf *os.File
        var iLen int
        var Ver byte

        if len(os.Args) != 2 {
                fmt.Printf("Please input mp3 filename\n")
                return
        }
        if path.Ext(os.Args[1]) != ".mp3" {
                fmt.Printf("Do not support this file format [%s]\n", path.Ext(os.Args[1]))
                return
        }
        if v, err := os.Open(os.Args[1]); err != nil {
                fmt.Printf("Error: %v\n", err)
        } else {
                rf = v
        }
        defer rf.Close()
        buffer := make([]byte, 10)
        if n, err := rf.Read(buffer); err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        } else {
                iLen = n
        }
        Ver = '1'
        if isAtHeadOfMP3(buffer[:iLen]) {
                Ver = '2'
        }
        fi, err := rf.Stat()
        if err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }
        if Ver == '2' {
                if buffer[3] != 0x03 {
                        fmt.Printf("Only support Version IDV3 Version 2.3\n")
                        return
                }
                iLen := uint32(buffer[6]&0x7f)<<21 | uint32(buffer[7]&0x7f)<<14 | uint32(buffer[8]&0x7f)<<7 | uint32(buffer[9]&0x7f) - 10
                bufferIn := make([]byte, int(iLen))
                if n, err := rf.Read(bufferIn); err != nil {
                        fmt.Printf("Error: %v\n", err)
                        return
                } else {
                        if n != len(bufferIn) {
                                fmt.Printf("Error: File is Bad\n")
                                return
                        }
                        i := 0
                        iCount := 0
                        // println("#####", n, "#####")
                        // k := 0
                        for i < n {
                                if bufferIn[i] == 0x0 {
                                        i++
                                        continue
                                }
                                j := i + 4
                                iCount = i
                                iLen := uint32(bufferIn[j])<<24 | uint32(bufferIn[j+1])<<16 | uint32(bufferIn[j+2])<<8 | uint32(bufferIn[j+3])
                                x := string(bufferIn[iCount : iCount+4])
                                tLen := int(iLen)
                                bufferTmp := bytes.TrimLeftFunc(bufferIn[i+10:i+10+tLen],
                                        func(r rune) bool { return r == 0x00 || r == 0x01 })
                                if bufferTmp == nil {
                                        i += 10 + tLen
                                        continue
                                }
                                switch x {
                                case "TIT2":
                                        outLen := 0
                                        bufferOut := make([]byte, len(bufferTmp)*4)
                                        if err := iconv.ConvertGBKToUTF8(bufferTmp, len(bufferTmp),
                                                bufferOut, &outLen); err != nil {
                                                byteCopy(bufferOut, bufferTmp, len(bufferTmp))
                                                outLen = len(bufferTmp)
                                        }
                                        fmt.Printf("标题: %s\n", string(bufferOut[:outLen]))
                                case "TPE1":
                                        outLen := 0
                                        bufferOut := make([]byte, len(bufferTmp)*4)
                                        if err := iconv.ConvertGBKToUTF8(bufferTmp, len(bufferTmp),
                                                bufferOut, &outLen); err != nil {
                                                byteCopy(bufferOut, bufferTmp, len(bufferTmp))
                                                outLen = len(bufferTmp)
                                        }
                                        fmt.Printf("作者: %s\n", string(bufferOut[:outLen]))
                                case "TALB":
                                        outLen := 0
                                        bufferOut := make([]byte, len(bufferTmp)*4)
                                        if err := iconv.ConvertGBKToUTF8(bufferTmp, len(bufferTmp),
                                                bufferOut, &outLen); err != nil {
                                                byteCopy(bufferOut, bufferTmp, len(bufferTmp))
                                                outLen = len(bufferTmp)
                                        }
                                        fmt.Printf("专辑名: %s\n", string(bufferOut[:outLen]))
                                case "TYER":
                                        outLen := 0
                                        bufferOut := make([]byte, len(bufferTmp)*4)
                                        if err := iconv.ConvertGBKToUTF8(bufferTmp, len(bufferTmp),
                                                bufferOut, &outLen); err != nil {
                                                byteCopy(bufferOut, bufferTmp, len(bufferTmp))
                                                outLen = len(bufferTmp)
                                        }
                                        fmt.Printf("发行年份: %s\n", string(bufferOut[:outLen]))
                                case "TCON":
                                        outLen := 0
                                        bufferOut := make([]byte, len(bufferTmp)*4)
                                        if err := iconv.ConvertGBKToUTF8(bufferTmp, len(bufferTmp),
                                                bufferOut, &outLen); err != nil {
                                                byteCopy(bufferOut, bufferTmp, len(bufferTmp))
                                                outLen = len(bufferTmp)
                                        }
                                        fmt.Printf("类型: %s\n", string(bufferOut[:outLen]))
                                case "COMM":
                                        outLen := 0
                                        bufferOut := make([]byte, len(bufferTmp)*4)
                                        if err := iconv.ConvertGBKToUTF8(bufferTmp, len(bufferTmp),
                                                bufferOut, &outLen); err != nil {
                                                byteCopy(bufferOut, bufferTmp, len(bufferTmp))
                                                outLen = len(bufferTmp)
                                        }
                                        fmt.Printf("备注信息: %s\n", string(bufferOut[:outLen]))
                                }
                                i += 10 + tLen
                                /*println(i, "|", iLen)
                                  k++
                                    if k == 25 {
                                            break
                                    }*/
                        }
                }

        } else {
                lPos := fi.Size() - 128
                bufferIn := make([]byte, 128)
                byteSet(bufferIn, 0, len(bufferIn))
                bufferOut := make([]byte, 1024)
                if n, err := rf.ReadAt(bufferIn, lPos); err != nil {
                        fmt.Printf("Error: %v\n", err)
                        return
                } else {
                        iLen = n
                }
                if !bytes.Equal(bufferIn[:3], []byte("TAG")) {
                        fmt.Printf("This file no idv3 information\n")
                        return
                }
                outLen := 0
                if err := iconv.ConvertGBKToUTF8(bufferIn[3:33], 30,
                        bufferOut, &outLen); err != nil {
                        fmt.Printf("Error: %v\n", err)
                        return
                }
                fmt.Printf("标题: %s\n", string(bufferOut[:outLen]))

                outLen = 0
                byteSet(bufferOut, 0, len(bufferOut))
                if err := iconv.ConvertGBKToUTF8(bufferIn[33:63], 30,
                        bufferOut, &outLen); err != nil {
                        fmt.Printf("Error: %v\n", err)
                        return
                }
                fmt.Printf("演唱/奏者: %s\n", string(bufferOut[:outLen]))

                outLen = 0
                byteSet(bufferOut, 0, len(bufferOut))
                if err := iconv.ConvertGBKToUTF8(bufferIn[63:93], 30,
                        bufferOut, &outLen); err != nil {
                        fmt.Printf("Error: %v\n", err)
                        return
                }
                fmt.Printf("专辑名: %s\n", string(bufferOut[:outLen]))

                outLen = 0
                byteSet(bufferOut, 0, len(bufferOut))
                if err := iconv.ConvertGBKToUTF8(bufferIn[93:97], 4,
                        bufferOut, &outLen); err != nil {
                        fmt.Printf("Error: %v\n", err)
                        return
                }
                fmt.Printf("发行年份: %s\n", string(bufferOut[:outLen]))

                outLen = 0
                byteSet(bufferOut, 0, len(bufferOut))
                if err := iconv.ConvertGBKToUTF8(bufferIn[97:127], 30,
                        bufferOut, &outLen); err != nil {
                        fmt.Printf("Error: %v\n", err)
                        return
                }
                fmt.Printf("备注信息: %s\n", string(bufferOut[:outLen]))
                if v, ok := LPInfo[bufferIn[127]]; !ok {
                        fmt.Printf("流派: %s\n", "未知")
                } else {
                        fmt.Printf("流派: %s\n", v)
                }
        }
}
