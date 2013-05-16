package main

import (
        "io"
        "log"
        "net/http"
)

func httpExampleFunc(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "This is index\n")
}

type myHttp struct{}

// hello world, the web server
func (c *myHttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, r.URL.Path)
        switch r.URL.Path {
        case "/":
                httpExampleFunc(w, r)
        default:
                /* func(w http.ResponseWriter, r *http.Request) {
                        io.WriteString(w, "Not found\n")
                }(w, r) */
                http.NotFound(w, r)
        }

}
func main() {
        // http.HandleFunc("/", ServerHandler)
        c := new(myHttp)
        err := http.ListenAndServe(":2525", c)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}
