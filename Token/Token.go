package token

import (
	"fmt"
	dbToken "myproject/DB/Token"
	"net/http"
)

// var validToken = "Bearer hdfjkgjkdfghi35jkhrgjndjdhgfjhdjfkgh34593405inasqweqwe823lkjk"

func CheckToken(r *http.Request, accountID int) (bool, error) {
	var headerToken = r.Header.Get("Authorization")
	fmt.Println("Header Token is ", headerToken)

	tokenOnDB, error := dbToken.FetchTokenByUserID(accountID)
	fmt.Println("Token on db is ", tokenOnDB)
	if error != nil {
		fmt.Println(error)
		return false, fmt.Errorf("user invalid")
	}

	if headerToken == "Bearer "+tokenOnDB {
		return true, nil
	}

	return false, fmt.Errorf("user invalid")
}
