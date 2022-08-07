package main

import (
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("./WWW")))
    http.ListenAndServe(":8080", nil)
}
