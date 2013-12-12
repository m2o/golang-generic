package main

import (
    "fmt"
    "net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s)
}

func main() {
    http.Handle("/burek",String("BUREK123"))
    http.ListenAndServe("localhost:4000", nil)
}

