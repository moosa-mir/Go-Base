package admin

import (
	"encoding/json"
	"net/http"
)

func (db *Admin) SellerListHandler(w http.ResponseWriter, r *http.Request) {
	sellers, err := db.DB.FetchSellers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(sellers)
}
