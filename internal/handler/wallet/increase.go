package wallet

import (
	"encoding/json"
	"fmt"
	"myproject/internal/auth"
	"myproject/internal/model"
	"net/http"
)

func (db *Api) HandleIncreaseWallet(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.FetchUserIDFromToken(r)
	if err != nil || userID == nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	var increaseModel model.IncreaseWalletModel
	errorJson := json.NewDecoder(r.Body).Decode(&increaseModel)
	if errorJson != nil {
		http.Error(w, errorJson.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("user is increasing wallet", increaseModel)
	errOnDB := db.DB.IncreaseWallet(increaseModel, *userID)
	if errOnDB != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
