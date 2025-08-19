package basket

import (
	"github.com/google/uuid"
	"time"
)

// Basket represents the core basket entity in the domain
type Basket struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Items     []BasketItem
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BasketItem represents an item in the basket
type BasketItem struct {
	ID        uuid.UUID
	BasketID  uuid.UUID
	ProductID uuid.UUID
	Quantity  int
	Price     float32
	SellerID  uuid.UUID
	CreatedAt time.Time
}

// NewBasket creates a new basket
func NewBasket(userID uuid.UUID) *Basket {
	return &Basket{
		ID:        uuid.New(),
		UserID:    userID,
		Items:     []BasketItem{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// NewBasketItem creates a new basket item
func NewBasketItem(productID uuid.UUID, quantity int, price float32, sellerID uuid.UUID) BasketItem {
	return BasketItem{
		ID:        uuid.New(),
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
		SellerID:  sellerID,
		CreatedAt: time.Now(),
	}
}

// AddItem adds an item to the basket
func (b *Basket) AddItem(item BasketItem) {
	// Check if item already exists
	for i, existingItem := range b.Items {
		if existingItem.ProductID == item.ProductID {
			// Update quantity
			b.Items[i].Quantity += item.Quantity
			b.UpdatedAt = time.Now()
			return
		}
	}
	
	// Add new item
	item.BasketID = b.ID
	b.Items = append(b.Items, item)
	b.UpdatedAt = time.Now()
}

// RemoveItem removes an item from the basket
func (b *Basket) RemoveItem(productID uuid.UUID) {
	for i, item := range b.Items {
		if item.ProductID == productID {
			b.Items = append(b.Items[:i], b.Items[i+1:]...)
			b.UpdatedAt = time.Now()
			break
		}
	}
}

// UpdateItemQuantity updates the quantity of an item
func (b *Basket) UpdateItemQuantity(productID uuid.UUID, quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	
	for i, item := range b.Items {
		if item.ProductID == productID {
			b.Items[i].Quantity = quantity
			b.UpdatedAt = time.Now()
			return nil
		}
	}
	
	return ErrItemNotFound
}

// GetTotal calculates the total price of all items
func (b *Basket) GetTotal() float32 {
	var total float32
	for _, item := range b.Items {
		total += item.Price * float32(item.Quantity)
	}
	return total
}

// IsEmpty checks if the basket is empty
func (b *Basket) IsEmpty() bool {
	return len(b.Items) == 0
}

// Clear removes all items from the basket
func (b *Basket) Clear() {
	b.Items = []BasketItem{}
	b.UpdatedAt = time.Now()
}

// Validate ensures the basket is valid
func (b *Basket) Validate() error {
	for _, item := range b.Items {
		if item.Quantity <= 0 {
			return ErrInvalidQuantity
		}
		if item.Price <= 0 {
			return ErrInvalidPrice
		}
	}
	return nil
}


