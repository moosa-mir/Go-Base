package token

import (
	"fmt"
	dbToken "myproject/DB/Token"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key-here")

type Claims struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func CheckToken(r *http.Request, accountID int) (bool, error) {
	var headerToken = r.Header.Get("Authorization")
	fmt.Println("Header Token is ", headerToken)
	fmt.Println("account id is ", accountID)

	tokenOnDB, error := dbToken.FetchTokenByUserID(accountID)
	fmt.Println("Token on db is ", tokenOnDB)
	if error != nil {
		fmt.Println(error)
		return false, fmt.Errorf("user invalid on fetchig token")
	}

	if headerToken == "Bearer "+tokenOnDB {
		return true, nil
	}
	
	return false, fmt.Errorf("user invalid for checking token")
}

func GenerateToken(userID int, name string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	claims := &Claims{
		UserID: userID,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "myproject",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
