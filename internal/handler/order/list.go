package order

import (
	"encoding/json"
	token "myproject/internal/auth"
	"net/http"
)

func (db *Order) ListOrdersHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := token.FetchUserIDFromToken(r)
	if err != nil || userID == nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	orders, err := db.DB.FetchOrdersByUserID(*userID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(orders)
	w.WriteHeader(http.StatusOK)
	return
}
