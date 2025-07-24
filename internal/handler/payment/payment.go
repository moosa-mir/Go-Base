package payment

import (
	"myproject/db"
)

type Payment struct {
	DB *db.DB
}

func NewPayment(db *db.DB) *Payment {
	return &Payment{DB: db}
}
