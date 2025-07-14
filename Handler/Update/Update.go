package update

import (
	"encoding/json"
	"fmt"
	dbUser "myproject/DB/User"
	user "myproject/User"
	"net/http"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
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

	userOnDB, errorFetch := dbUser.FetchUserByUsername(userName)
	if userOnDB == nil || errorFetch != nil {
		fmt.Println(errorFetch)
		http.Error(w, "username is not valid on db", http.StatusBadRequest)
		return
	}

	result := dbUser.UpdateUser(updateModel, userName)
	if !result {
		http.Error(w, "Internal error", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "User updated successfully"}`)
}
