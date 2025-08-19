package basket

import (
	"github.com/google/uuid"
)

// Service contains business logic for basket operations
type Service struct {
	repo Repository
}

// NewService creates a new basket service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetBasket retrieves a basket by user ID
func (s *Service) GetBasket(userID uuid.UUID) (*Basket, error) {
	return s.repo.FindByUserID(userID)
}

// AddItemToBasket adds an item to the user's basket
func (s *Service) AddItemToBasket(userID, productID uuid.UUID, quantity int, price float32, sellerID uuid.UUID) error {
	// Check if product is already in basket
	exists, err := s.repo.IsProductInBasket(productID, userID)
	if err != nil {
		return err
	}
	
	if exists {
		// Update existing item quantity
		basket, err := s.repo.FindByUserID(userID)
		if err != nil {
			return err
		}
		
		if err := basket.UpdateItemQuantity(productID, quantity); err != nil {
			return err
		}
		
		return s.repo.Update(basket)
	}
	
	// Add new item
	item := NewBasketItem(productID, quantity, price, sellerID)
	return s.repo.AddItem(userID, item)
}

// RemoveItemFromBasket removes an item from the user's basket
func (s *Service) RemoveItemFromBasket(userID, productID uuid.UUID) error {
	return s.repo.RemoveItem(userID, productID)
}

// ClearBasket clears all items from the user's basket
func (s *Service) ClearBasket(userID uuid.UUID) error {
	return s.repo.DeleteByUserID(userID)
}

// GetBasketTotal calculates the total price of items in the basket
func (s *Service) GetBasketTotal(userID uuid.UUID) (float32, error) {
	basket, err := s.repo.FindByUserID(userID)
	if err != nil {
		return 0, err
	}
	
	return basket.GetTotal(), nil
}

// IsProductInBasket checks if a product is in the user's basket
func (s *Service) IsProductInBasket(productID, userID uuid.UUID) (bool, error) {
	return s.repo.IsProductInBasket(productID, userID)
}


