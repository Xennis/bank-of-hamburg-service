package main

import "time"

var accountCurrentId uint
var accounts map[uint]Account
var transactions Transactions
var transactionCurrentId uint

func init() {
	accounts = make(map[uint]Account)

	DbCreateAccount(Account{Customer: "Tommy Schmidt"})
	DbCreateAccount(Account{Customer: "Frida Sandberg"})
}

func DbFindAccount(id uint) Account {
	a, ok := accounts[id]
	if ok == true {
		return a
	}

	// Found no account => Return empty object
	return Account{}
}

func DbCreateAccount(a Account) Account {
	accountCurrentId += 1 // Super safe way ...
	a.Id = accountCurrentId
	a.Cash = 0
	a.Created = time.Now()
	a.Verified = true // Verify all accounts by default
	accounts[a.Id] = a
	return a
}

func DbUpdateAccount(id uint, cash float64) Account {
	a := DbFindAccount(id)
	// TODO: Take care, if no account was found
	a.Cash += cash
	accounts[id] = a
	return a
}

func DbCreateTransaction(t Transaction) Transaction {
	t.Id = transactionCurrentId
	if t.From > 0 {
		DbUpdateAccount(t.From, -t.Amount)
	}
	if t.To > 0 {
		DbUpdateAccount(t.To, t.Amount)
	}
	t.Created = time.Now()
	transactions = append(transactions, t)
	return t
}
