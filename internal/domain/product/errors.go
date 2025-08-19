package product

import "errors"

var (
	ErrInvalidName        = errors.New("invalid product name")
	ErrInvalidPrice       = errors.New("invalid product price")
	ErrInvalidStock       = errors.New("invalid product stock")
	ErrInsufficientStock  = errors.New("insufficient stock")
	ErrProductNotFound    = errors.New("product not found")
)

