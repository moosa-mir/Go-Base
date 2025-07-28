package account

import (
	"encoding/json"
	"fmt"
	token "myproject/internal/auth"
	"myproject/internal/model"
	"myproject/internal/utils"
	"net/http"
)

// Define a struct to represent the login response
type TokenResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

func (db *Account) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body into a LoginRequest struct
	var loginReq model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println("Login Request is ", loginReq)
	// Validate the credentials
	userOnDB, errDB := db.DB.FetchUserByUsername(loginReq.Username)
	if userOnDB == nil || errDB != nil {
		fmt.Println("user is not valid ", loginReq)
		http.Error(w, "Invalid data", http.StatusUnauthorized)
		return
	}

	hashedPassword := utils.GetHashPassword(loginReq.Password)
	fmt.Println("Hash password is", hashedPassword)
	if hashedPassword == "" {
		http.Error(w, "Invalid data", http.StatusUnauthorized)
		return
	}

	if userOnDB.Password != hashedPassword {
		http.Error(w, "Invalid data", http.StatusUnauthorized)
		return
	}

	newToken, err := token.GenerateToken(userOnDB.Username, userOnDB.ID)
	if newToken == "" || err != nil {
		http.Error(w, "Internal Error Generate Token", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := TokenResponse{
		Token: newToken,
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)

}
