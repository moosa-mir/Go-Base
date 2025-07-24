package account

import (
	"encoding/json"
	"fmt"
	token "myproject/internal/auth"
	user "myproject/internal/model"
	_ "myproject/internal/utils"
	"net/http"
)

func (db *Account) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not support", http.StatusMethodNotAllowed)
		return
	}

	userID, err := token.FetchUserIDFromToken(r)
	if err != nil || userID == nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
	}
	var updateModel user.UpdateUser
	errDecode := json.NewDecoder(r.Body).Decode(&updateModel)
	if errDecode != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userOnDB, errorFetch := db.DB.FetchUserByUserID(*userID)
	if userOnDB == nil || errorFetch != nil {
		fmt.Println(errorFetch)
		http.Error(w, "username is not valid on db", http.StatusBadRequest)
		return
	}

	result, _ := db.DB.UpdateUser(updateModel, *userID)
	if !result {
		http.Error(w, "Internal error", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "User updated successfully"}`)
}
