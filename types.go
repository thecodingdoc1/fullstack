package main

import (
	"errors"
	"math/rand"
)

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int64
	Balance   int64
}

var Accounts = make(map[string]*Account)

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(100000)),
	}
}

func AddAccount(account *Account) error {
	key := account.LastName + "," + account.FirstName
	// Check if the key already exists
	if _, exists := Accounts[key]; exists {
		return errors.New("account with this name already exists")
	}
	Accounts[key] = account

	return nil

}
