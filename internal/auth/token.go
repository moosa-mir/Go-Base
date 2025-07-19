package auth

import (
	"myproject/internal/constant"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Phone    string `json:"phone"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string, phone string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	claims := &Claims{
		Username: username,
		Phone:    phone,
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

func FetchUsernameFromToken(r *http.Request) (string, error) {
	tokenString := FetchToken(r)
	// Parse the token with the secret key
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		// Return the secret key used to sign the token
		return constant.JwtSecret, nil
	})

	if err != nil {
		// Handle token parsing errors (e.g., invalid token, expired token)
		return "", err
	}

	// Check if the token is valid and has claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Extract and return the username from the claims
		return claims.Username, nil
	}

	// If the token is invalid or claims are missing
	return "", jwt.ErrInvalidKey
}

func FetchToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	tokenString := authHeader // assuming format: Bearer <token>
	if tokenString == "Bearer " {
		return ""
	}

	// Remove "Bearer " prefix if needed
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	return tokenString
}
