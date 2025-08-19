package seller

import (
	"github.com/google/uuid"
)

// Repository defines the contract for seller data access
type Repository interface {
	FindByID(id uuid.UUID) (*Seller, error)
	FindByUsername(username string) (*Seller, error)
	FindAll() ([]*Seller, error)
	Save(seller *Seller) error
	Update(seller *Seller) error
	Exists(username string) (bool, error)
}


