package auth

import (
	"fmt"
	"log"
	"myproject/internal/constant"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	UserKey   contextKey = "user"
	ClaimsKey contextKey = "claims"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := FetchToken(r)

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
