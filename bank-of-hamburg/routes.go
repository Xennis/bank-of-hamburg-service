package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "AccountIndex",
        "GET",
        "/accounts",
        AccountIndex,
    },
    Route{
        "AccountCreate",
        "POST",
        "/accounts",
        AccountCreate,
    },
    Route{
        "AccountShow",
        "GET",
        "/accounts/{accountId}",
        AccountShow,
    },
    Route{
        "TransactionCreate",
        "POST",
        "/transactions",
        TransactionCreate,
    },
}
