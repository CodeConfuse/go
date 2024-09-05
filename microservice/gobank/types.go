package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID        int       `json:"id`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, LastName string) *Account {

	return &Account{
		//	ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  LastName,
		Number:    int64(rand.Intn(100000)),
		CreatedAt: time.Now(),
	}
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
