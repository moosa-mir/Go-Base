package account

import (
	"encoding/json"
	"fmt"
	user "myproject/internal/model"
	"myproject/internal/utils"
	"net/http"
)

func (db *Api) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not support", http.StatusMethodNotAllowed)
		return
	}

	username := utils.GetQueryParam(r, "username")
	var updateModel user.UpdateUser
	err := json.NewDecoder(r.Body).Decode(&updateModel)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userOnDB, errorFetch := db.DB.FetchUserByUsername(username)
	if userOnDB == nil || errorFetch != nil {
		fmt.Println(errorFetch)
		http.Error(w, "username is not valid on db", http.StatusBadRequest)
		return
	}

	result, _ := db.DB.UpdateUser(updateModel, username)
	if !result {
		http.Error(w, "Internal error", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "User updated successfully"}`)
}
