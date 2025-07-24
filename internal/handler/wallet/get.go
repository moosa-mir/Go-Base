package wallet

import (
	"encoding/json"
	"fmt"
	"myproject/internal/auth"
	"myproject/internal/model"
	"net/http"
)

func (db *Api) WalletHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.FetchUserIDFromToken(r)
	if err != nil || userID == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	walletModel, err := db.DB.FetchWalletByUserID(*userID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if walletModel == nil {
		_ = json.NewEncoder(w).Encode(model.WalletModel{})
		return
	}
	_ = json.NewEncoder(w).Encode(walletModel)
	w.WriteHeader(http.StatusOK)
}
