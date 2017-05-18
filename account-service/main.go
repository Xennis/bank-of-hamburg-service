package main

import (
    "log"
    "flag"
    "net/http"


)

const (
	addr = ":8080" // own rest service
)

var (
	transactionServiceAddr = flag.String("addr", "transaction-service:80", "Address of transaction service")
	transactionServiceKey  = flag.String("api-key", "", "API key.")
)


func main() {
    flag.Parse()

    router := NewRouter()
    log.Fatal(http.ListenAndServe(addr, router))
}
