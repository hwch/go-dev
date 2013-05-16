package main

import (
        "fmt"
        "net"
)

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
        raddr, err := net.ResolveUDPAddr("udp", "192.168.135.254:520")
        if err != nil {
                fmt.Printf("Error1: %v", err)
                return
        }
        /* laddr, err1 := net.ResolveUDPAddr("udp", ":520")
           if err1 != nil {
                   fmt.Printf("Error2: %v", err1)
                   return
           } */

        ufd, err2 := net.DialUDP("udp", nil, raddr)
        if err2 != nil {
                fmt.Printf("Error3: %v", err2)
                return
        }
        defer ufd.Close()
        sendBuffer := make([]byte, 512)
        sendBuffer[0] = 1 // 命令
        sendBuffer[1] = 1 // 版本
        sendBuffer[5] = 0 // 地址系列 ?????? 4,5
        byteCopy(sendBuffer[8:], net.ParseIP("192.168.135.254"), 4)
        sendBuffer[23] = 16 //度量值 ??????  20,23

        if _, err := ufd.Write(sendBuffer[:24]); err != nil {
                fmt.Printf("Error4: %v", err)
                return
        }
        fmt.Printf("Send success\n")
        byteSet(sendBuffer, 0x00, len(sendBuffer))
        if n, _, err := ufd.ReadFromUDP(sendBuffer); err != nil {
                fmt.Printf("Error5: %v", err)
                return
        } else {
                fmt.Printf("%v\n", sendBuffer[:n])
        }
        /*lu, err3 := net.ListenUDP("udp", laddr)
          if err3 != nil {
                  fmt.Printf("Error6: %v", err3)
                  return
          }
          if n, _, err := lu.ReadFrom(sendBuffer); err != nil {
                  fmt.Printf("Error5: %v", err)
                  return
          } else {
                  fmt.Printf("%v\n", sendBuffer[:n])
          } */
}
