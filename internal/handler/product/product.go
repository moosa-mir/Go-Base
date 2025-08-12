package product

import (
	"myproject/db"
)

type Product struct {
	DB db.ProductDBInterface
}

func NewProduct(db db.ProductDBInterface) *Product {
	return &Product{DB: db}
}
