package product

import (
	"github.com/google/uuid"
)

// Repository defines the contract for product data access
type Repository interface {
	FindByID(id uuid.UUID) (*Product, error)
	FindAll() ([]*Product, error)
	FindBySeller(sellerID uuid.UUID) ([]*Product, error)
	Search(keyword string) ([]*Product, error)
	Save(product *Product) error
	Update(product *Product) error
	Delete(id uuid.UUID) error
}

