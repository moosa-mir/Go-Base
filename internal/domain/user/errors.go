package user

import "errors"

var (
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidName     = errors.New("invalid name")
	ErrUserNotFound    = errors.New("user not found")
	ErrUserExists      = errors.New("user already exists")
)

