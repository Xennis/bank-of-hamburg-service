package main

import (
    "log"
    "net/http"
)

const (
	addr = ":8080" // Adress of the service
)

func main() {
    router := NewRouter()
    log.Fatal(http.ListenAndServe(addr, router))
}
