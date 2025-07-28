package seller

import (
	json "encoding/json"
	"fmt"
	model "myproject/internal/model"
	http "net/http"
)

func (db *Seller) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("RegisterSellerHandler")

	// Decode the JSON body into a LoginRequest struct
	var registerModel model.RegistrationSeller
	err := json.NewDecoder(r.Body).Decode(&registerModel)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userOnDB, err := db.DB.FetchSellerByUsername(registerModel.Username)
	if err == nil || userOnDB != nil {
		http.Error(w, "seller exist", http.StatusConflict)
		return
	}

	result, _ := db.DB.InsertSeller(registerModel)
	if result {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"message": "User registered successfully"}`)
		return
	} else {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
}
