package main

import (
        "fmt"
        "os"
)

func startDaemon() error {
        var proc *os.Process

        dir, err := os.Getwd()
        if err != nil {
                return err
        }
        procattr := os.ProcAttr{Dir: dir, Env: os.Environ(), Files: []*os.File{nil, nil, nil}}
        proc, err = os.StartProcess("helloworld.exe", nil, &procattr)
        if err != nil {
                return err
        }
        return proc.Release()
}

func main() {
        if err := startDaemon(); err != nil {
                fmt.Printf("Error:%v\n", err)
                return
        }
        fmt.Printf("Daemon start sucess ...\n")
}
