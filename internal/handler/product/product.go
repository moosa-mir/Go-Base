package product

import (
	"myproject/db"
)

type Product struct {
	DB *db.DB
}

func NewProduct(db *db.DB) *Product {
	return &Product{DB: db}
}
