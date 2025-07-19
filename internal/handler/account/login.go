package account

import (
	"encoding/json"
	"fmt"
	token "myproject/internal/auth"
	"net/http"
)

// Define a struct to represent the login request payload
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Define a struct to represent the login response
type TokenResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

func (db *Api) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body into a LoginRequest struct
	var loginReq LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the credentials
	userOnDB, error := db.DB.FetchUserByUsername(loginReq.Username)
	// userOnDB, error := dbUser.FetchUserByUsername(loginReq.Username)
	if userOnDB == nil || error != nil {
		fmt.Println("user is not valid ", loginReq)
		http.Error(w, "Invalid data", http.StatusUnauthorized)
		return
	}

	newToken, error := token.GenerateToken(userOnDB.Username, string(userOnDB.Phone))
	if newToken == "" || error != nil {
		http.Error(w, "Internal Error Generate Token", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := TokenResponse{
		Token: newToken,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
