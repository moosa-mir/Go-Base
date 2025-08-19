package user

import (
	"github.com/google/uuid"
	"myproject/internal/utils"
)

// Service contains business logic for user operations
type Service struct {
	repo Repository
}

// NewService creates a new user service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// RegisterUser registers a new user
func (s *Service) RegisterUser(username, password, name, family string, birthday string, city, country, phone string) error {
	// Parse birthday
	birthdayTime, err := utils.ParseDate(birthday)
	if err != nil {
		return err
	}

	// Create user entity
	user := NewUser(username, password, name, family, birthdayTime, city, country, phone)

	// Validate user
	if err := user.Validate(); err != nil {
		return err
	}

	// Check if user already exists
	exists, err := s.repo.Exists(username)
	if err != nil {
		return err
	}
	if exists {
		return ErrUserExists
	}

	// Hash password
	user.Password = utils.GetHashPassword(password)

	// Save user
	return s.repo.Save(user)
}

// GetUserByID retrieves a user by ID
func (s *Service) GetUserByID(id uuid.UUID) (*User, error) {
	return s.repo.FindByID(id)
}

// GetUserByUsername retrieves a user by username
func (s *Service) GetUserByUsername(username string) (*User, error) {
	return s.repo.FindByUsername(username)
}

// UpdateUserProfile updates a user's profile
func (s *Service) UpdateUserProfile(id uuid.UUID, name, family string) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	user.UpdateProfile(name, family)
	return s.repo.Update(user)
}

