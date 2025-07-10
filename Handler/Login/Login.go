package login

import (
	"encoding/json"
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

// Simulated database of users (for demonstration purposes)
var validUser = LoginRequest{
	Username: "admin",
	Password: "password123",
}

// LoginHandler handles the login logic
func LoginHandler(w http.ResponseWriter, r *http.Request) {
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
	if loginReq.Username == validUser.Username && loginReq.Password == validUser.Password {
		// Successful login
		response := TokenResponse{
			Token: "hdfjkgjkdfghi35jkhrgjndjdhgfjhdjfkgh34593405inasqweqwe823lkjk",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		// Failed login
		response := ErrorResponse{
			Message:   "Invalid username or password",
			ErrorCode: 401,
		}
		w.WriteHeader(http.StatusUnauthorized) // 401 Unauthorized
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
