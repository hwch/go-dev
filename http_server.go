package main

import (
        "fmt"
        "html/template"
        "log"
        "net/http"
        "strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        // 解析url 传递的参数,对于POST 则解析响应包的主体(request body )
        // 注意: 如果没有调用ParseForm 方法,下面无法获取表单的数据
        fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
        fmt.Println("path", r.URL.Path)
        fmt.Println("scheme", r.URL.Scheme)
        fmt.Println(r.Form["url_long"])
        for k, v := range r.Form {
                fmt.Println("key:", k)
                fmt.Println("val:", strings.Join(v, ""))
        }
        fmt.Fprintf(w, "Hello hwch!") // 这个写入到w 的是输出到客户端的
}
func login(w http.ResponseWriter, r *http.Request) {

        fmt.Println("method:", r.Method) // 获取请求的方法
        if r.Method == "GET" {
                t, _ := template.ParseFiles("login.tmp")
                t.Execute(w, nil)
        } else {
                if err := r.ParseForm(); err != nil {
                        panic("ParseForm")
                }
                // 请求的是登陆数据,那么执行登陆的逻辑判断
                fmt.Println("username:", r.Form["username"])
                fmt.Println("password:", r.Form["password"])
        }
}

func main() {
        http.HandleFunc("/", sayhelloName)
        // 设置访问的路由
        http.HandleFunc("/login", login)
        // 设置访问的路由
        err := http.ListenAndServe(":2525", nil) // 设置监听的端口
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}
