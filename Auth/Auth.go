package auth

import (
	"fmt"
	"log"
	constant "myproject/Constant"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	UserKey   contextKey = "user"
	ClaimsKey contextKey = "claims"
)

type Claims struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := authHeader // assuming format: Bearer <token>
		if tokenString == "Bearer " {
			http.Error(w, "Invalid or missing token", http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " prefix if needed
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return constant.JwtSecret, nil
		})

		if err != nil {
			log.Println("Token parsing error:", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Token is valid → proceed
		next(w, r)
	}
}
