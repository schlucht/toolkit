package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type Account struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		Balance:   0,
		CreatedAt: time.Now().UTC(),
	}
}
