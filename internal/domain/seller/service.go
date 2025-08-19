package seller

import (
	"myproject/internal/utils"
)

// Service contains business logic for seller operations
type Service struct {
	repo Repository
}

// NewService creates a new seller service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// RegisterSeller registers a new seller
func (s *Service) RegisterSeller(username, password, name, email, phone, address string) error {
	// Check if seller already exists
	exists, err := s.repo.Exists(username)
	if err != nil {
		return err
	}
	if exists {
		return ErrSellerExists
	}

	// Create seller entity
	seller := NewSeller(username, password, name, email, phone, address)

	// Validate seller
	if err := seller.Validate(); err != nil {
		return err
	}

	// Hash password
	seller.Password = utils.GetHashPassword(password)

	// Save seller
	return s.repo.Save(seller)
}

// GetSellerByUsername retrieves a seller by username
func (s *Service) GetSellerByUsername(username string) (*Seller, error) {
	return s.repo.FindByUsername(username)
}

// GetSellerByID retrieves a seller by ID
func (s *Service) GetSellerByID(id uuid.UUID) (*Seller, error) {
	return s.repo.FindByID(id)
}

// GetAllSellers retrieves all sellers
func (s *Service) GetAllSellers() ([]*Seller, error) {
	return s.repo.FindAll()
}

// UpdateSellerProfile updates a seller's profile
func (s *Service) UpdateSellerProfile(id uuid.UUID, name, email, phone, address string) error {
	seller, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	seller.UpdateProfile(name, email, phone, address)
	return s.repo.Update(seller)
}

// UpdateSellerStatus updates a seller's status
func (s *Service) UpdateSellerStatus(id uuid.UUID, status SellerStatus) error {
	seller, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	seller.UpdateStatus(status)
	return s.repo.Update(seller)
}


