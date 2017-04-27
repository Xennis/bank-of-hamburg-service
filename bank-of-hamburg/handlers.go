package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)


const (
	contentType = "Content-Type"
	contentTypeJson = "application/json; charset=UTF-8"
)

// Show all accounts (no pagination)
func AccountIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentTypeJson)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		panic(err)
	}
}

// Create a new account
func AccountCreate(w http.ResponseWriter, r *http.Request) {
	var account Account
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &account); err != nil {
		w.Header().Set(contentType, contentTypeJson)
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	p := DbCreateAccount(account)
	w.Header().Set(contentType, contentTypeJson)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}
}

// Show details of a account
func AccountShow(w http.ResponseWriter, r *http.Request) {
    const paramId string = "accountId"

	vars := mux.Vars(r)
	accountIdUint64, err := strconv.ParseUint(vars[paramId], 0, 32);
	if err != nil {
		panic(err)
	}
	accountId := uint(accountIdUint64)
	account := DbFindAccount(accountId)
	if account.Id > 0 {
		w.Header().Set(contentType, contentTypeJson)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(account); err != nil {
			panic(err)
		}
		return
	}

	// Found no account => Return error code
	w.Header().Set(contentType, contentTypeJson)
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}


// Create a new transaction
func TransactionCreate(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &transaction); err != nil {
		w.Header().Set(contentType, contentTypeJson)
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// TODO: Check from or to is valid
	// TODO: Check amount is positive

	p := DbCreateTransaction(transaction)
	w.Header().Set(contentType, contentTypeJson)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}
}
