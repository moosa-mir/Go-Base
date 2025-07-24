package account

import (
	json "encoding/json"
	"fmt"
	_ "github.com/google/uuid"
	token "myproject/internal/auth"
	http "net/http"
)

func (db *Account) AccountInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is Get
	fmt.Println("AccountInfoHandler")

	userID, errorUserID := token.FetchUserIDFromToken(r)
	if errorUserID != nil || userID == nil {
		http.Error(w, "Internal Error Fetch Username From Token", http.StatusUnauthorized)
		return
	}
	user, errorDB := db.DB.FetchUserByUserID(*userID)
	if errorDB != nil || user == nil {
		fmt.Println(errorDB)
		http.Error(w, "Invalid Account", http.StatusBadRequest)
		return
	}

	result := user.ConvertToUser()
	fmt.Println("username is", userID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
