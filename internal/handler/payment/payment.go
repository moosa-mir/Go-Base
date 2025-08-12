package payment

import (
	"myproject/db"
)

type Payment struct {
	DB db.PaymentDBInterface
}

func NewPayment(db db.PaymentDBInterface) *Payment {
	return &Payment{DB: db}
}
