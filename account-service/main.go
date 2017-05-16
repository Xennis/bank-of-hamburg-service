package main

import (
    "log"
    "net/http"


)

const (
	addr = ":8080" // own rest service
    transactionApiAddr = "localhost:50051" // grpc api of the transaction service
)


func main() {
    router := NewRouter()
    log.Fatal(http.ListenAndServe(addr, router))
}
