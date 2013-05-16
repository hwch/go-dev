package main

import (
        //"bufio"
        "fmt"
        "io/ioutil"
        "net/rpc"
        "os"
        "runtime"
        "strconv"
        "strings"
        "time"
)

const MTU_LEN = 4096

func Test_panic() {
        /* if x := recover(); x != nil {
        	println("Panic")
        	// return
        } */
        pQuit <- true

}

type MyTestIn struct {
        Args0 int
}
type MyTestOut struct {
        Args0 string
}
type GetFileInSt struct {
        FileName string
        Offset   int64
        MaxUnit  int64
}

type GetFileOutSt struct {
        TotalSize int64
        DataLen   int64
        Data      []byte
}

var pQuit chan bool

func ToConnection() {
        var in MyTestIn
        var out MyTestOut
        // fmt.Printf("进程ID为:%d\n", os.Getpid())
        conn, err := rpc.Dial("tcp", "localhost:52525")
        defer Test_panic()
        if err != nil {
                // handle error
                fmt.Printf("Dial error: %v\n", err)
                // panic("TimeOut")
                return
        }
        defer conn.Close()

        // fmt.Fprintf(conn, "Hello World\n")
        // status, err := bufio.NewReader(conn).ReadString('\n')
        for i := 0; i < 100; i++ {
                in.Args0 = 10000 + i
                if err := conn.Call("RPC.RPCMyTest", &in, &out); err != nil {
                        fmt.Printf("ERROR:%v\n", err)
                        return
                }
                fmt.Printf("Recieve: %s\n", out.Args0)
                time.Sleep(time.Second)
        }
}

func RemoteProcessCall(iTimes int) error {
        var FileName string
        var MaxUnit int64
        var Offset int64
        var TotalSize int64

        if iTimes == 0 {
                return nil
        }

        conn, err := rpc.Dial("tcp", "localhost:12525")
        defer Test_panic()
        if err != nil {
                // handle error
                fmt.Printf("Dial error: %v\n", err)
                // panic("TimeOut")
                return err
        }
        defer conn.Close()

        MaxUnit = MTU_LEN
        FileName = "123.txt"
        DownLoadFile := FileName + ".hdownload"

        if _, err := os.Stat(DownLoadFile); err != nil {
                Offset = 0
                if f, err := os.Create(DownLoadFile); err != nil {
                        return err
                } else {
                        defer f.Close()
                        if _, err := f.WriteString("0:0"); err != nil {
                                return err
                        }

                }
        } else {
                if data, err := ioutil.ReadFile(DownLoadFile); err != nil {
                        return err
                } else {
                        if data[0] == 'A' {
                                return nil
                        } else {
                                x := strings.Split(string(data), ":")
                                TotalSize, err = strconv.ParseInt(x[0], 10, 64)
                                if err != nil {
                                        return err
                                }
                                Offset, err = strconv.ParseInt(x[1], 10, 64)
                                if TotalSize != 0 && Offset == TotalSize {
                                        return nil
                                }
                        }
                }
        }
        if _, err := os.Stat(FileName); err != nil {
                if f, err := os.Create(FileName); err != nil {
                        return err
                } else {
                        f.Close()
                }
        }

        iCount := 0
        f, verr := os.OpenFile(FileName, os.O_RDWR, 0644)
        if verr != nil {
                return verr
        }
        defer f.Close()
        for {
                if iTimes > 0 && iCount >= iTimes {
                        break
                }
                in := new(GetFileInSt)
                out := new(GetFileOutSt)
                in.Offset = Offset
                in.MaxUnit = MaxUnit
                in.FileName = FileName
                if err := conn.Call("RPC.RPC_GetFile", &in, &out); err != nil {
                        fmt.Printf("ERROR:%v\n", err)
                        return err
                }
                Offset = in.Offset + out.DataLen
                TotalSize = out.TotalSize
                if in.Offset+out.DataLen == out.TotalSize {
                        _, err := f.WriteAt(out.Data[:out.DataLen], in.Offset)
                        if err != nil {
                                return err
                        }
                        break
                }
                _, err := f.WriteAt(out.Data[:out.DataLen], in.Offset)
                if err != nil {
                        return err
                }
                Offset = in.Offset + out.DataLen
                TotalSize = out.TotalSize
                iCount++
        }

        myStr := fmt.Sprintf("%d:%d", TotalSize, Offset)
        if err := ioutil.WriteFile(DownLoadFile, []byte(myStr), 0644); err != nil {
                os.Remove(DownLoadFile)
                return err
        }

        return nil
}

func ToConnection1() {
        var in MyTestIn
        var out MyTestOut
        // fmt.Printf("进程ID为:%d\n", os.Getpid())
        conn, err := rpc.Dial("tcp", "localhost:12525")
        defer Test_panic()
        if err != nil {
                // handle error
                fmt.Printf("Dial error: %v\n", err)
                // panic("TimeOut")
                return
        }
        defer conn.Close()

        // fmt.Fprintf(conn, "Hello World\n")
        // status, err := bufio.NewReader(conn).ReadString('\n')
        for i := 0; i < 100; i++ {
                in.Args0 = 10000 + i
                if err := conn.Call("RPC.MyTest", &in, &out); err != nil {
                        fmt.Printf("ERROR:%v\n", err)
                        //return
                }

                fmt.Printf("Recieve1: %s\n", out.Args0)
                time.Sleep(time.Second)
        }
}

func main() {
        pQuit = make(chan bool)
        defer close(pQuit)
        runtime.GOMAXPROCS(2)

        /* go ToConnection() */
        go RemoteProcessCall(1)

        <-pQuit
        // fmt.Printf("进程ID为:%d\n", os.Getpid())
}
