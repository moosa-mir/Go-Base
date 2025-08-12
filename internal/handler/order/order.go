package order

import "myproject/db"

type Order struct {
	DB db.OrderDBInterface
}

func NewOrder(db db.OrderDBInterface) *Order {
	return &Order{DB: db}
}
