package order

import "myproject/db"

type Order struct {
	DB *db.DB
}

func NewOrder(db *db.DB) *Order {
	return &Order{DB: db}
}
