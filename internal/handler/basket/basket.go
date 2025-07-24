package basket

import (
	"myproject/db"
)

type Basket struct {
	DB *db.DB
}

func NewBasket(db *db.DB) *Basket {
	return &Basket{DB: db}
}
