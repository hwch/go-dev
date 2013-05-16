package main

import (
        "fmt"
        "reflect"
)

func PrintType(v interface{}) {
        i := -1
        tp := []string{"byte[uint8]", "int8", "uint16", "int16",
                "uint32", "int32", "uint64", "int64", "int", "uint",
                "bool", "uintptr", "float32", "float64", "complex64",
                "complex128", "arry", "chan", "func", "interface",
                "map", "ptr", "clice", "string", "struct",
                "unsafePointer", "unkwon type"}
        switch t := reflect.TypeOf(v); t.Kind() {
        case reflect.Uint8:
                i = 0
        case reflect.Int8:
                i = 1
        case reflect.Uint16:
                i = 2
        case reflect.Int16:
                i = 3
        case reflect.Uint32:
                i = 4
        case reflect.Int32:
                i = 5
        case reflect.Uint64:
                i = 6
        case reflect.Int64:
                i = 7
        case reflect.Int:
                i = 8
        case reflect.Uint:
                i = 9
        case reflect.Bool:
                i = 10
        case reflect.Uintptr:
                i = 11
        case reflect.Float32:
                i = 12
        case reflect.Float64:
                i = 13
        case reflect.Complex64:
                i = 14
        case reflect.Complex128:
                i = 15
        case reflect.Array:
                i = 16
        case reflect.Chan:
                i = 17
        case reflect.Func:
                i = 18
        case reflect.Interface:
                i = 19
        case reflect.Map:
                i = 20
        case reflect.Ptr:
                i = 21
        case reflect.Slice:
                i = 22
        case reflect.String:
                i = 23
        case reflect.Struct:
                i = 24
        case reflect.UnsafePointer:
                i = 25
        default:
                i = 26
        }

        fmt.Printf("Type: %s\n", tp[i])
}

func main() {
        var a byte
        var b int8
        var c uint16
        var d int16
        var e uint32
        var f int32
        var g uint64
        var h int64
        var i int
        var j uint
        fu := func(int) int {
                return 0
        }

        PrintType(a)
        PrintType(b)
        PrintType(c)
        PrintType(d)
        PrintType(e)
        PrintType(f)
        PrintType(g)
        PrintType(h)
        PrintType(i)
        PrintType(j)
        PrintType(fu)
}
