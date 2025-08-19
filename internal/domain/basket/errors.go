package basket

import "errors"

var (
	ErrInvalidQuantity = errors.New("invalid quantity")
	ErrInvalidPrice    = errors.New("invalid price")
	ErrItemNotFound    = errors.New("item not found")
	ErrBasketNotFound  = errors.New("basket not found")
	ErrEmptyBasket     = errors.New("basket is empty")
)


