package main

import "time"

type Transaction struct {
    Id          uint        `json:"id"`
    From        uint        `json:"from"`
    To          uint        `json:"to"`
    Amount      float64     `json:"amount"`
    Created     time.Time   `json:"created"`
}

type Transactions []Transaction
