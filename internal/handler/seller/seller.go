package seller

import "myproject/db"

type Seller struct {
	DB *db.DB
}

func NewSeller(db *db.DB) *Seller {
	return &Seller{DB: db}
}
