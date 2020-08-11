package main

import (
        "flag"
        "fmt"
        "net/http"
)

var s string

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, s)
}

func main() {
        var str = flag.String("str", "Hello World", "string flag")
        flag.Parse()
        s = *str
        http.HandleFunc("/", handler)
        http.ListenAndServe(":8080", nil)
}
