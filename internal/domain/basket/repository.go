package basket

import (
	"github.com/google/uuid"
)

// Repository defines the contract for basket data access
type Repository interface {
	FindByUserID(userID uuid.UUID) (*Basket, error)
	Save(basket *Basket) error
	Update(basket *Basket) error
	DeleteByUserID(userID uuid.UUID) error
	AddItem(userID uuid.UUID, item BasketItem) error
	RemoveItem(userID uuid.UUID, productID uuid.UUID) error
	IsProductInBasket(productID, userID uuid.UUID) (bool, error)
}


