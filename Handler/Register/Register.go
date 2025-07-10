package register

import (
	json "encoding/json"
	"fmt"
	dbUser "myproject/DB/User"
	user "myproject/User"
	http "net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
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

	userOnDB, error := dbUser.FetchUserByUsername(registerModel.Username)
	if error == nil || userOnDB != nil {
		http.Error(w, "username exist", http.StatusConflict)
		return
	}
	dbUser.Init()
	result := dbUser.InsertUser(registerModel)
	if result {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"message": "User registered successfully"}`)
		return
	} else {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
}
