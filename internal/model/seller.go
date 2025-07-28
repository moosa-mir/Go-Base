package model

import "github.com/google/uuid"

type Seller struct {
	ID            uuid.UUID `json:"id"`
	Username      string    `json:"username"`
	Name          string    `json:"name"`
	Family        string    `json:"family"`
	Phone         country   `json:"phone"`
	AccountNumber country   `json:"account_number"`
	Address       string    `json:"address"`
}

type StoredSeller struct {
	ID            uuid.UUID `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Name          string    `json:"name"`
	Family        string    `json:"family"`
	Phone         country   `json:"phone"`
	AccountNumber country   `json:"account_number"`
	Address       string    `json:"address"`
}

type SellerStored interface {
	ConvertToSeller() User
}

func (u StoredSeller) ConvertToSeller() Seller {
	seller := Seller{Username: u.Username, Name: u.Name, Family: u.Family, Phone: u.Phone, AccountNumber: u.AccountNumber, Address: u.Address}
	return seller
}

type RegistrationSeller struct {
	Username      string  `json:"username"`
	Password      string  `json:"password"`
	Name          string  `json:"name"`
	Family        string  `json:"family"`
	Phone         country `json:"phone"`
	AccountNumber country `json:"account_number"`
	Address       string  `json:"address"`
}
