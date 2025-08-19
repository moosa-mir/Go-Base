package product

import (
	"github.com/google/uuid"
)

// Product represents the core product entity in the domain
type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       float32
	SellerID    uuid.UUID
	Category    string
	ImageURL    string
	Stock       int
}

// NewProduct creates a new product entity
func NewProduct(name, description string, price float32, sellerID uuid.UUID, category, imageURL string, stock int) *Product {
	return &Product{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
		SellerID:    sellerID,
		Category:    category,
		ImageURL:    imageURL,
		Stock:       stock,
	}
}

// UpdateStock updates the product stock
func (p *Product) UpdateStock(newStock int) {
	p.Stock = newStock
}

// DecreaseStock decreases stock by the given amount
func (p *Product) DecreaseStock(amount int) error {
	if p.Stock < amount {
		return ErrInsufficientStock
	}
	p.Stock -= amount
	return nil
}

// Validate ensures the product entity is valid
func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrInvalidName
	}
	if p.Price <= 0 {
		return ErrInvalidPrice
	}
	if p.Stock < 0 {
		return ErrInvalidStock
	}
	return nil
}

