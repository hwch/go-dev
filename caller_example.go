package main

import (
        "fmt"
        "runtime"
)

func getFuncInfo() (funcName, fileName string, lineNo int) {
        var pc uintptr
        pc, fileName, lineNo, _ = runtime.Caller(1)
        pc, _, _, _ = runtime.Caller(1)
        f := runtime.FuncForPC(pc)
        funcName = f.Name()
        return
}

func main() {
        a, b, c := getFuncInfo()

        fmt.Printf("%s:%s:%d\n", a, b, c)
}
