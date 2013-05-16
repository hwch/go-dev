package main

import (
        "fmt"
        "os"
        "path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
        fmt.Printf("FileName[%s]FileSize[%d]Dir[%v]\n",
                filepath.Join(path, info.Name()), info.Size(), info.IsDir())
        return err
}

func main() {
        if err := filepath.Walk("E:\\go-dev\\atmc", walkFunc); err != nil {
                fmt.Printf("Error: %v\n", err)
        } else {
                fmt.Println("Complete!!!!!!!!!")
        }
}
