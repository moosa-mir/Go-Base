package wallet

import (
	"github.com/google/uuid"
)

// Service contains business logic for wallet operations
type Service struct {
	repo Repository
}

// NewService creates a new wallet service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetWallet retrieves a wallet by user ID
func (s *Service) GetWallet(userID uuid.UUID) (*Wallet, error) {
	return s.repo.FindByUserID(userID)
}

// CreateWallet creates a new wallet for a user
func (s *Service) CreateWallet(userID uuid.UUID) (*Wallet, error) {
	wallet := NewWallet(userID)
	
	if err := wallet.Validate(); err != nil {
		return nil, err
	}
	
	if err := s.repo.Save(wallet); err != nil {
		return nil, err
	}
	
	return wallet, nil
}

// IncreaseBalance increases the wallet balance
func (s *Service) IncreaseBalance(userID uuid.UUID, amount float32) error {
	wallet, err := s.repo.FindByUserID(userID)
	if err != nil {
		// If wallet doesn't exist, create one
		if err == ErrWalletNotFound {
			wallet = NewWallet(userID)
		} else {
			return err
		}
	}
	
	if err := wallet.IncreaseBalance(amount); err != nil {
		return err
	}
	
	if wallet.ID == uuid.Nil {
		return s.repo.Save(wallet)
	}
	
	return s.repo.Update(wallet)
}

// DecreaseBalance decreases the wallet balance
func (s *Service) DecreaseBalance(userID uuid.UUID, amount float32) error {
	wallet, err := s.repo.FindByUserID(userID)
	if err != nil {
		return err
	}
	
	if err := wallet.DecreaseBalance(amount); err != nil {
		return err
	}
	
	return s.repo.Update(wallet)
}

// GetBalance returns the current balance of a user's wallet
func (s *Service) GetBalance(userID uuid.UUID) (float32, error) {
	wallet, err := s.repo.FindByUserID(userID)
	if err != nil {
		return 0, err
	}
	
	return wallet.GetBalance(), nil
}


