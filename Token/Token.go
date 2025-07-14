package token

import (
	auth "myproject/Auth"
	constant "myproject/Constant"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int, name string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	claims := &auth.Claims{
		UserID: userID,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "myproject",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(constant.JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
