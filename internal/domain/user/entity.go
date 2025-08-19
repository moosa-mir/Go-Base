package user

import (
	"github.com/google/uuid"
	"time"
)

// User represents the core user entity in the domain
type User struct {
	ID       uuid.UUID
	Username string
	Password string
	Name     string
	Family   string
	Birthday time.Time
	City     string
	Country  string
	Phone    string
}

// NewUser creates a new user entity
func NewUser(username, password, name, family string, birthday time.Time, city, country, phone string) *User {
	return &User{
		ID:       uuid.New(),
		Username: username,
		Password: password,
		Name:     name,
		Family:   family,
		Birthday: birthday,
		City:     city,
		Country:  country,
		Phone:    phone,
	}
}

// UpdateProfile updates the user's profile information
func (u *User) UpdateProfile(name, family string) {
	u.Name = name
	u.Family = family
}

// Validate ensures the user entity is valid
func (u *User) Validate() error {
	if u.Username == "" {
		return ErrInvalidUsername
	}
	if u.Password == "" {
		return ErrInvalidPassword
	}
	if u.Name == "" {
		return ErrInvalidName
	}
	return nil
}

