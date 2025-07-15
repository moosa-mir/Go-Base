package login

import (
	"encoding/json"
	db "myproject/DB/Init"
	dbToken "myproject/DB/Token"
	dbUser "myproject/DB/User"
	token "myproject/Token"
	"net/http"
)

// Define a struct to represent the login request payload
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Define a struct to represent the login response
type TokenResponse struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}

type LoginApi struct {
	DB *db.DB
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

func LoginApiHandler(db *db.DB) *LoginApi {
	return &LoginApi{DB: db}
}

func (db *LoginApi) LoginHandler(w http.ResponseWriter, r *http.Request) {
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
	userOnDB, error := dbUser.FetchUserByUsername(loginReq.Username)
	if userOnDB == nil || error != nil {
		http.Error(w, "Invalid data", http.StatusUnauthorized)
		return
	}

	newToken, error := token.GenerateToken(userOnDB.UserID, userOnDB.Username)
	if newToken == "" || error != nil {
		http.Error(w, "Internal Error Generate Token", http.StatusConflict)
		return
	}

	result := dbToken.InsertTokenForUserID(userOnDB.UserID, newToken)
	if !result {
		http.Error(w, "Internal Error Insert Token", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := TokenResponse{
		Token:  newToken,
		UserID: userOnDB.UserID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	// Ensure the request method is POST
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Decode the JSON body into a LoginRequest struct
// 	var loginReq LoginRequest
// 	err := json.NewDecoder(r.Body).Decode(&loginReq)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	// Validate the credentials
// 	userOnDB, error := dbUser.FetchUserByUsername(loginReq.Username)
// 	if userOnDB == nil || error != nil {
// 		http.Error(w, "Invalid data", http.StatusUnauthorized)
// 		return
// 	}

// 	newToken, error := token.GenerateToken(userOnDB.UserID, userOnDB.Username)
// 	if newToken == "" || error != nil {
// 		http.Error(w, "Internal Error Generate Token", http.StatusConflict)
// 		return
// 	}

// 	result := dbToken.InsertTokenForUserID(userOnDB.UserID, newToken)
// 	if !result {
// 		http.Error(w, "Internal Error Insert Token", http.StatusConflict)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := TokenResponse{
// 		Token:  newToken,
// 		UserID: userOnDB.UserID,
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }
