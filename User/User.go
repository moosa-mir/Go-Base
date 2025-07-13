package user

type User struct {
	UserID   int     `json:"user_id"`
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Family   string  `json:"family"`
	Birthday int     `json:"birthday"`
	City     city    `json:"city"`
	Country  country `json:"country"`
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
	UserID   int     `json:"user_id"`
	Name     string  `json:"name"`
	Family   string  `json:"family"`
	Birthday int     `json:"birthday"`
	City     city    `json:"city"`
	Country  country `json:"country"`
	Username string  `json:"username"`
	Password string  `json:"password"`
}

type UserStoredExtention interface {
	convertToUse() User
}

func (u StoredUser) ConvertToUser() User {
	user := User{UserID: u.UserID, Username: u.Username, Name: u.Name, Family: u.Family, Birthday: u.Birthday, City: u.City, Country: u.Country}
	return user
}

// ///////////////////////////////
type RegistrationUser struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	Family   string  `json:"family"`
	Birthday int     `json:"birthday"`
	City     city    `json:"city"`
	Country  country `json:"country"`
}

// ///////////////////////////////
type UserExtention interface {
	calculateAge() int
	giveFullName() string
	updateCity(newCity city)
}

// Implement calculateAge
func (u User) calculateAge() int {
	return 2023 - u.Birthday
}

// Implement giveFullName
func (u User) giveFullName() string {
	return u.Name + " " + u.Family
}

func (u *User) updateCity(newCity city) {
	u.City = newCity
}

type UpdateUser struct {
	Name   string `json:"name"`
	Family string `json:"family"`
}
