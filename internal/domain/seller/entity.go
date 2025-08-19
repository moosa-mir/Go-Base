package seller

import (
	"github.com/google/uuid"
	"time"
)

// Seller represents the core seller entity in the domain
type Seller struct {
	ID        uuid.UUID
	Username  string
	Password  string
	Name      string
	Email     string
	Phone     string
	Address   string
	Status    SellerStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

// SellerStatus represents the status of a seller
type SellerStatus string

const (
	SellerStatusActive   SellerStatus = "active"
	SellerStatusInactive SellerStatus = "inactive"
	SellerStatusSuspended SellerStatus = "suspended"
)

// NewSeller creates a new seller
func NewSeller(username, password, name, email, phone, address string) *Seller {
	return &Seller{
		ID:        uuid.New(),
		Username:  username,
		Password:  password,
		Name:      name,
		Email:     email,
		Phone:     phone,
		Address:   address,
		Status:    SellerStatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// UpdateProfile updates the seller's profile information
func (s *Seller) UpdateProfile(name, email, phone, address string) {
	s.Name = name
	s.Email = email
	s.Phone = phone
	s.Address = address
	s.UpdatedAt = time.Now()
}

// UpdateStatus updates the seller's status
func (s *Seller) UpdateStatus(status SellerStatus) {
	s.Status = status
	s.UpdatedAt = time.Now()
}

// IsActive checks if the seller is active
func (s *Seller) IsActive() bool {
	return s.Status == SellerStatusActive
}

// Validate ensures the seller entity is valid
func (s *Seller) Validate() error {
	if s.Username == "" {
		return ErrInvalidUsername
	}
	if s.Password == "" {
		return ErrInvalidPassword
	}
	if s.Name == "" {
		return ErrInvalidName
	}
	if s.Email == "" {
		return ErrInvalidEmail
	}
	return nil
}


