package accountinfo

import (
	json "encoding/json"
	"fmt"
	db "myproject/DB"
	utils "myproject/Utils"
	http "net/http"
)

type AccountApi struct {
	DB *db.DB
}

func ApiHandler(db *db.DB) *AccountApi {
	return &AccountApi{DB: db}
}

func (db *AccountApi) AccountInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is Get
	fmt.Println("AccountInfoHandler")
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Extract the path parameter "username"
	username, error := utils.GetPathString(r, 2)
	if error != nil {
		fmt.Println("No account ID found")
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	user, error := db.DB.FetchUserByUsername(username)
	if error != nil || user == nil {
		fmt.Println(error)
		http.Error(w, "Invalid Account", http.StatusBadRequest)
		return
	}

	result := user.ConvertToUser()
	fmt.Println("username is", username)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
