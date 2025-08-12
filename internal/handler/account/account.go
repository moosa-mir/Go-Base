package account

import (
	"myproject/db"
)

type Account struct {
	DB db.UserDBInterface
}

func NewAccount(db db.UserDBInterface) *Account {
	return &Account{DB: db}
}
