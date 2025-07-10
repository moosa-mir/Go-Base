package main

import (
	registerRoute "myproject/RegisterRoute"
)

func main() {
	// dbUser.Init()
	// userRegistration := user.RegistrationUser{Username: "moosamir", Password: "0000", Name: "Moosa", Family: "Mir", Birthday: 1371, City: user.Tehran, Country: user.Iran}
	// dbUser.InsertUser(userRegistration)

	// dbToken.Init()
	// token, error := dbToken.FetchTokenByUserID(2)
	// fmt.Println("token for 2 is ", token)
	// if error != nil {
	// 	dbToken.InsertTokenForUserID(2, "hdfjkgjkdfghi35jkhrgjndjdhgfjhdjfkgh34593405inasqweqwe823lkjk")
	// }

	registerRoute.RegisterRoutes()

}
