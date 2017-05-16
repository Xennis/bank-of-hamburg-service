package main

import (
    "log"
    "os"
    "google.golang.org/grpc"
)

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

func getTransactionClient() TransactionApiClient {
    conn, err := grpc.Dial(transactionApiAddr, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    return NewTransactionApiClient(conn)
}
