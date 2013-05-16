package main

/*
#include <stdio.h>
#include <stdlib.h>
FILE *getStdout(void) { return stdout; }
*/
import "C"
import "unsafe"

func main() {
        x := C.CString("Hello World\n")
        defer C.free(unsafe.Pointer(x))
        C.fputs(x, C.getStdout())

}
