package admin

import (
	"encoding/json"
	"net/http"
)

func (db *Admin) FetchUnPayoutHandler(w http.ResponseWriter, r *http.Request) {
	unPayoutOrders, err := db.DB.FetchUnPayoutOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(unPayoutOrders)
	w.WriteHeader(http.StatusOK)
}
