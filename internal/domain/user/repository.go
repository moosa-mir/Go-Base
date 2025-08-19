package user

import (
	"github.com/google/uuid"
)

// Repository defines the contract for user data access
type Repository interface {
	FindByID(id uuid.UUID) (*User, error)
	FindByUsername(username string) (*User, error)
	Save(user *User) error
	Update(user *User) error
	Exists(username string) (bool, error)
}

