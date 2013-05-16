package main

import (
        "errors"
        "fmt"
        "os"
        "strconv"
        "sync"
        "time"
)

const (
        RPT_TO_FILE  uint    = 0X10
        RPT_TO_STOUT uint    = 0X20
)

const (
        OFF_LEVEL  uint    = 0x0000
        ERR_LEVEL  uint    = 0x1000
        WARN_LEVEL uint    = 0x2000
        INFO_LEVEL uint    = 0x3000
)

const DEFAULT_LOG_NAME = "atmcDebug.log"

type AtmLogSt struct {
        logName    string
        debugLevel uint
        logLevel   [3]string
        rwLock     *sync.RWMutex
}

// 初始化日志相关参数
func InitLog(logName string, debugSwitch uint) *AtmLogSt {
        als := new(AtmLogSt)
        als.logName = logName
        als.logLevel[ERR_LEVEL>>12-1] = "ERR"
        als.logLevel[INFO_LEVEL>>12-1] = "INFO"
        als.logLevel[WARN_LEVEL>>12-1] = "WARN"
        als.rwLock = new(sync.RWMutex)
        als.debugLevel = debugSwitch

        return als
}

// 更改打印日志级别
func (als *AtmLogSt) ChgLogLevel(debugSwitch uint) {
        als.debugLevel = debugSwitch
}

// 将相关信息写入日志文件
func (als *AtmLogSt) WriteLog(LogLevel uint, FileName string,
        FuncName string, LineNo int,
        LogMsg string, LogOutObj uint,
        DumpMsg []byte, DumpLen int) error {

        var err error
        var LFile *os.File
        var WStr string
        //++++++++++++++++++++++++++++++++++++++++++++++++
        defer func() {
                myc <- 1
        }()
        //++++++++++++++++++++++++++++++++++++++++++++++++
        if LogLevel > als.debugLevel {
                return nil
        }

        switch LogLevel {
        case ERR_LEVEL:
                fallthrough
        case WARN_LEVEL:
                fallthrough
        case INFO_LEVEL:
        default:
                return errors.New("不可识别的日志级别!")
        }
        Ltime := time.Now()

        WStr = WStr + fmt.Sprintf("Pid [%d] | Time [%02d:%02d:%02d] | Message: \n",
                os.Getpid(), Ltime.Hour(), Ltime.Minute(), Ltime.Second())

        WStr = WStr + fmt.Sprintf("%s>> File[%s] Function[%s] Line[%d] %s |\n",
                als.logLevel[LogLevel>>12-1], FileName, FuncName, LineNo, LogMsg)
        if RPT_TO_FILE^LogOutObj == 0 {
                LFile, err = os.OpenFile(als.logName, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_EXCL, 0644)
                if err != nil {
                        if os.IsExist(err) {
                                LFile, err = os.OpenFile(als.logName, os.O_APPEND|os.O_RDWR, 0644)
                                if err != nil {
                                        return err
                                }
                        } else {
                                return err
                        }
                }
                defer LFile.Close()
        } else {
                LFile = os.Stdout
        }

        if DumpMsg != nil && DumpLen > 0 {
                if tStr, err := als.memLogWrite(DumpMsg, DumpLen); err != nil {
                        return err
                } else {
                        WStr = WStr + tStr
                }
        }

        als.rwLock.Lock()
        LFile.WriteString(WStr)
        als.rwLock.Unlock()

        return nil
}

// 以ASCII码形式对照打印原始字符串
func (als *AtmLogSt) memLogWrite(DumpMsg []byte, DumpLen int) (string, error) {
        Count := 0
        Pos := 0

        WStr := "Displacement    ++00++01++02++03++04++05++06++07++08++09++10++11++12++13++" +
                "14++15++  ++ASCII  Value++\n"
        Len := len(WStr)

L:
        for DumpLen != 0 {
                WStr = WStr + fmt.Sprintf("%05d(%05X)      ",
                        Count, Count)
                Len += 18 /* length of strlen("%05d(%05x)     ++"); */
                if DumpLen < 16 {
                        TmpBuf := make([]byte, DumpLen)
                        copy(TmpBuf, DumpMsg[Pos:Pos+DumpLen])
                        i := 0
                        for i < DumpLen {
                                WStr = WStr + fmt.Sprintf("%02X  ", TmpBuf[i])
                                Len += 4
                                i++
                        }
                        for i < 16 {
                                WStr = WStr + "    "
                                Len += 4
                                i++
                        }
                        WStr = WStr + "  "
                        Len += 2
                        if err := toPrintCh(TmpBuf, DumpLen); err != nil {
                                return "", err
                        }
                        WStr = WStr + string(TmpBuf) + "\n"
                        Len += DumpLen + 1
                        break L
                } else {
                        TmpBuf := make([]byte, 16)
                        copy(TmpBuf, DumpMsg[Pos:Pos+16])
                        DumpLen -= 16
                        WStr = WStr + fmt.Sprintf("%02X  %02X  %02X  %02X  "+
                                "%02X  %02X  %02X  %02X  "+
                                "%02X  %02X  %02X  %02X  "+
                                "%02X  %02X  %02X  %02X    ",
                                TmpBuf[0], TmpBuf[1], TmpBuf[2], TmpBuf[3],
                                TmpBuf[4], TmpBuf[5], TmpBuf[6], TmpBuf[7],
                                TmpBuf[8], TmpBuf[9], TmpBuf[10], TmpBuf[11],
                                TmpBuf[12], TmpBuf[13], TmpBuf[14], TmpBuf[15])
                        Len += 4*16 + 2
                        if err := toPrintCh(TmpBuf, 16); err != nil {
                                return "", err
                        }
                        WStr = WStr + string(TmpBuf) + "\n"
                        Len += 17
                        Pos += 16
                }
                Count += 16
        }

        return WStr, nil
}

// 将数字0转换为可见字符'*'
func toPrintCh(v []byte, slen int) error {
        if v == nil {
                return errors.New("输入参数不能为空")
        }

        for i := 0; i < slen; i++ {
                if v[i] == 0 {
                        v[i] = '*'
                }
        }
        return nil
}

var myc chan int

func main() {
        y := "大家好才是真的好，如果那么。有志者事竟成，总有城要有要民圾加蝇浊极婚。"
        z := "This is a test"
        x := InitLog("test.log", INFO_LEVEL)
        myc = make(chan int)
        defer close(myc)

        for i := 1; i < 100; i++ {
                t := strconv.Itoa(i)
                go x.WriteLog(INFO_LEVEL, "test.log", "main", 193, z+t, RPT_TO_FILE, []byte(y+t), len(y+t))
        }

LL:
        for {
                select {
                case <-myc:
                case <-time.After(time.Second * 2):
                        break LL
                }
        }

}
