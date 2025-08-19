package seller

import "errors"

var (
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidName     = errors.New("invalid name")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrSellerNotFound  = errors.New("seller not found")
	ErrSellerExists    = errors.New("seller already exists")
	ErrInactiveSeller  = errors.New("seller is inactive")
)


