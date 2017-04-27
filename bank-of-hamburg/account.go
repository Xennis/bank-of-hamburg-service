package main

import "time"

type Account struct {
    Id          uint        `json:"id"`
    Customer    string      `json:"customer"`
    Created     time.Time   `json:"created"`
    Verified    bool        `json:"verified"`
    Cash        float64     `json:"cash"`
}

type Accounts []Account
