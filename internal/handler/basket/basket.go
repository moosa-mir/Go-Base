package basket

import (
	"myproject/db"
)

type Basket struct {
	DB db.BasketDBInterface
}

func NewBasket(db db.BasketDBInterface) *Basket {
	return &Basket{DB: db}
}
