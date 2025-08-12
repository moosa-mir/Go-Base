package seller

import "myproject/db"

type Seller struct {
	DB db.SellerDBInterface
}

func NewSeller(db db.SellerDBInterface) *Seller {
	return &Seller{DB: db}
}
