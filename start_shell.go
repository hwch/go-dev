package main

import "os"
import "fmt"

func main() {
        if p, err := os.StartProcess("/bin/bash", []string{"/bin/bash", "-l"},
                &os.ProcAttr{Files: []*os.File{os.Stdin,
                        os.Stdout, os.Stderr}}); err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        } else {
                p.Wait()
                fmt.Printf("/bin/bash run success!!!\n")
        }
}
