package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

const (
	api_prefix = "/api" // Adress of the service
)

var routes = Routes{
    Route{
        "LivenessProbe",
        "GET",
        "/healthz",
        LivenessProbe,
    },
    Route{
        "ReadinessProbe",
        "GET",
        "/readiness",
        ReadinessProbe,
    },
    Route{
        "AccountIndex",
        "GET",
        api_prefix + "/accounts",
        AccountIndex,
    },
    Route{
        "AccountCreate",
        "POST",
        api_prefix + "/accounts",
        AccountCreate,
    },
    Route{
        "AccountShow",
        "GET",
        api_prefix + "/accounts/{accountId}",
        AccountShow,
    },
    Route{
        "TransactionCreate",
        "POST",
        api_prefix + "/transactions",
        TransactionCreate,
    },
}
