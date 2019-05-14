package main

import (
    "fmt"
    "net/http"
    "os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    host := os.Getenv("HOSTNAME")
    fmt.Fprintln(w, "hello world "+ host +", this is version 1." )
}


func main()  {
    http.Handle("/pattern", http.HandlerFunc(IndexHandler))
    http.ListenAndServe(":80", nil)
}
