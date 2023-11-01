package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	Name string `json:"name"`
}

type Account struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Number    int64     `json:"number"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(name string) *Account {
	return &Account{
		Name:    name,
		Number:  rand.Int63(),
		Balance: 0,
	}
}
