package main

import (
    "fmt"
    "net/http"
)

func books(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "[\"haha\", \"two\"]")
}

func main() {
    http.HandleFunc("/book", books)
    fs := http.FileServer(http.Dir("go/portal/build"))
    http.Handle("/", fs)
    http.ListenAndServe(":8090", nil)
}
