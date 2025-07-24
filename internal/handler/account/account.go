package account

import (
	"myproject/db"
)

type Account struct {
	DB *db.DB
}

func NewAccount(db *db.DB) *Account {
	return &Account{DB: db}
}
