package model

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	Image       string    `json:"image"`
}
