package wallet

import (
	"encoding/json"
	token "myproject/internal/auth"
	"net/http"
)

func (db *Api) HandleWalletItems(w http.ResponseWriter, r *http.Request) {
	username, errorUsername := token.FetchUsernameFromToken(r)
	if username == "" || errorUsername != nil {
		http.Error(w, "User name is not valid", http.StatusConflict)
		return
	}
	items, err := db.DB.FetchWalletByUsername(username)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusConflict)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
