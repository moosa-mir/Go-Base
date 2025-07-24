package account

import (
	json "encoding/json"
	"fmt"
	user "myproject/internal/model"
	http "net/http"
)

func (db *Account) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	// func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is Get
	fmt.Println("RegisterHandler")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body into a LoginRequest struct
	var registerModel user.RegistrationUser
	err := json.NewDecoder(r.Body).Decode(&registerModel)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userOnDB, error := db.DB.FetchUserByUsername(registerModel.Username)
	if error == nil || userOnDB != nil {
		http.Error(w, "username exist", http.StatusConflict)
		return
	}

	result, _ := db.DB.InsertUser(registerModel)
	if result {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"message": "User registered successfully"}`)
		return
	} else {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
}
