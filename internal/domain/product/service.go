package product

import (
	"github.com/google/uuid"
	"strings"
)

// Service contains business logic for product operations
type Service struct {
	repo Repository
}

// NewService creates a new product service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetProduct retrieves a product by ID
func (s *Service) GetProduct(id uuid.UUID) (*Product, error) {
	return s.repo.FindByID(id)
}

// GetAllProducts retrieves all products
func (s *Service) GetAllProducts() ([]*Product, error) {
	return s.repo.FindAll()
}

// SearchProducts searches for products by keyword
func (s *Service) SearchProducts(keyword string) ([]*Product, error) {
	if strings.TrimSpace(keyword) == "" {
		return s.repo.FindAll()
	}
	return s.repo.Search(keyword)
}

// GetProductsBySeller retrieves products by seller ID
func (s *Service) GetProductsBySeller(sellerID uuid.UUID) ([]*Product, error) {
	return s.repo.FindBySeller(sellerID)
}

// CreateProduct creates a new product
func (s *Service) CreateProduct(name, description string, price float32, sellerID uuid.UUID, category, imageURL string, stock int) error {
	product := NewProduct(name, description, price, sellerID, category, imageURL, stock)
	
	if err := product.Validate(); err != nil {
		return err
	}
	
	return s.repo.Save(product)
}

// UpdateProductStock updates product stock
func (s *Service) UpdateProductStock(id uuid.UUID, newStock int) error {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	
	product.UpdateStock(newStock)
	return s.repo.Update(product)
}

