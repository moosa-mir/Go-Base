package account

import (
	"encoding/json"
	"fmt"
	db "myproject/DB"
	user "myproject/User"
	"net/http"
)

type AccountApi struct {
	DB *db.DB
}

func ApiHandler(db *db.DB) *AccountApi {
	return &AccountApi{DB: db}
}

func (db *AccountApi) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not support", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	userName := query.Get("username")
	if userName == "" {
		http.Error(w, "username is not valid on param", http.StatusBadRequest)
		return
	}

	var updateModel user.UpdateUser
	err := json.NewDecoder(r.Body).Decode(&updateModel)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userOnDB, errorFetch := db.DB.FetchUserByUsername(userName)
	if userOnDB == nil || errorFetch != nil {
		fmt.Println(errorFetch)
		http.Error(w, "username is not valid on db", http.StatusBadRequest)
		return
	}

	result, _ := db.DB.UpdateUser(updateModel, userName)
	if !result {
		http.Error(w, "Internal error", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "User updated successfully"}`)
}
