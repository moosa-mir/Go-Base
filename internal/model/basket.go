package model

import "github.com/google/uuid"

type StoredBasketItem struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	Date      string    `json:"date"`
	Count     int       `json:"count"`
}

type InsertBasketItem struct {
	ProductID uuid.UUID `json:"product_id"`
	Date      float32   `json:"date"`
	Count     int       `json:"count"`
}

type InputAddBasketItem struct {
	ProductID uuid.UUID `json:"product_id"`
}

type InputDeleteBasketItem struct {
	ProductID uuid.UUID `json:"product_id"`
}

type BasketItem struct {
	ID      uuid.UUID `json:"id"`
	Product Product   `json:"product"`
	Date    string    `json:"date"`
	Count   int       `json:"count"`
}
