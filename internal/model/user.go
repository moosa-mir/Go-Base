package model

import "github.com/google/uuid"

type User struct {
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Family   string  `json:"family"`
	Birthday int     `json:"birthday"`
	City     city    `json:"city"`
	Country  country `json:"country"`
	Phone    country `json:"phone"`
}

type city string

const (
	Tehran city = "Tehran"
)

type country string

const (
	Iran   country = "Iran"
	US     country = "US"
	AU     country = "AU"
	China  country = "China"
	Franch country = "Franch"
)

// ////////////////
type StoredUser struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Family   string    `json:"family"`
	Birthday int       `json:"birthday"`
	City     city      `json:"city"`
	Country  country   `json:"country"`
	Phone    country   `json:"phone"`
}

type UserStored interface {
	convertToUse() User
}

func (u StoredUser) ConvertToUser() User {
	user := User{Username: u.Username, Name: u.Name, Family: u.Family, Birthday: u.Birthday, City: u.City, Country: u.Country, Phone: u.Phone}
	return user
}

type RegistrationUser struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	Family   string  `json:"family"`
	Birthday int     `json:"birthday"`
	City     city    `json:"city"`
	Country  country `json:"country"`
	Phone    country `json:"phone"`
}

type UpdateUser struct {
	Name   string `json:"name"`
	Family string `json:"family"`
}
