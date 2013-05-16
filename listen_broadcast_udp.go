package main

import (
        "fmt"
        "hlog"
        "iconv"
        "io"
        "net"
)

// 设置字节
func byteSet(src []byte, s byte, iLen int) {
        for i := 0; i < iLen; i++ {
                src[i] = s
        }
}

type Logger interface {
        WriteLog(uint, string, uint, []byte, int) error
}

func main() {

        tcp_chan := make(chan bool)
        udp_chan := make(chan bool)

        // 初始化日志打印级别
        hlog.InitLogLevel(hlog.DEBUG_LEVEL)
        // 初始化日志
        myLog := hlog.InitLog("FQ.log")
        myLog.ChgLogFuncStyle(hlog.SHORT_FUNC)
        go func(myLog Logger, udp_chan chan bool) {
                defer func(c chan bool) {
                        c <- true
                }(udp_chan)
                uc, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4zero, Port: 2425})
                if err != nil {
                        fmt.Printf("Error:%v\n", err)
                        return
                }
                defer uc.Close()
                myLog.WriteLog(hlog.DEBUG_LEVEL, "UDP Listen 2425 success",
                        hlog.RPT_TO_FILE, nil, 0)

                buffer := make([]byte, 4096)
                recvBuffer := make([]byte, 8192)
                for {
                        n, raddr, err := uc.ReadFrom(buffer)
                        if err != nil {
                                myLog.WriteLog(hlog.ERR_LEVEL, err.Error(),
                                        hlog.RPT_TO_FILE, nil, 0)
                                break
                        }
                        iLen := 0
                        if err := iconv.ConvertGBKToUTF8(buffer, n, recvBuffer, &iLen); err != nil {
                                myLog.WriteLog(hlog.ERR_LEVEL, err.Error(),
                                        hlog.RPT_TO_FILE, buffer, n)
                                break
                        }
                        myLog.WriteLog(hlog.DEBUG_LEVEL, fmt.Sprintf("UDP收到%v", raddr),
                                hlog.RPT_TO_FILE, recvBuffer, iLen)
                        byteSet(buffer, 0, len(buffer))
                        byteSet(recvBuffer, 0, len(recvBuffer))
                }
        }(myLog, udp_chan)

        go func(myLog Logger, tcp_chan chan bool) {
                defer func(c chan bool) {
                        c <- true
                }(tcp_chan)
                tc, err := net.ListenTCP("tcp4", &net.TCPAddr{IP: net.IPv4zero, Port: 2425})
                if err != nil {
                        fmt.Printf("Error:%v\n", err)
                        return
                }
                defer tc.Close()
                myLog.WriteLog(hlog.DEBUG_LEVEL, "TCP Listen 2425 success",
                        hlog.RPT_TO_FILE, nil, 0)

                for {
                        if client, err := tc.AcceptTCP(); err != nil {
                                myLog.WriteLog(hlog.ERR_LEVEL, err.Error(),
                                        hlog.RPT_TO_FILE, nil, 0)
                                break
                        } else {
                                go func(cln *net.TCPConn) {
                                        defer cln.Close()
                                        buffer := make([]byte, 4096)
                                        for {
                                                n, err := cln.Read(buffer)
                                                if err != nil {
                                                        if err == io.EOF {
                                                                myLog.WriteLog(hlog.ERR_LEVEL, "对方关闭了连接",
                                                                        hlog.RPT_TO_FILE, nil, 0)
                                                        } else {
                                                                myLog.WriteLog(hlog.ERR_LEVEL, err.Error(),
                                                                        hlog.RPT_TO_FILE, nil, 0)
                                                        }
                                                        break
                                                }
                                                myLog.WriteLog(hlog.DEBUG_LEVEL, fmt.Sprintf("TCP收到%v", cln.RemoteAddr()),
                                                        hlog.RPT_TO_FILE, buffer, n)
                                                if _, err := cln.Write(buffer[:n]); err != nil {
                                                        myLog.WriteLog(hlog.ERR_LEVEL, err.Error(),
                                                                hlog.RPT_TO_FILE, nil, 0)
                                                        break

                                                }
                                                byteSet(buffer, 0, len(buffer))
                                        }
                                }(client)
                        }
                }
        }(myLog, tcp_chan)
        iCount := 0

        for iCount < 2 {
                select {
                case c := <-udp_chan:
                        if c {
                                myLog.WriteLog(hlog.ERR_LEVEL, "UDP Listen quit", hlog.RPT_TO_FILE, nil, 0)
                                iCount++
                        }
                case c := <-tcp_chan:
                        if c {
                                myLog.WriteLog(hlog.ERR_LEVEL, "TCP Listen quit", hlog.RPT_TO_FILE, nil, 0)
                                iCount++
                        }
                }
        }
}
