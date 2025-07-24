package payment

import (
	"fmt"
	token "myproject/internal/auth"
	"myproject/internal/model"
	"net/http"
)

func (db *Payment) PayHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := token.FetchUserIDFromToken(r)
	if err != nil || userID == nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	basket, err := db.DB.FetchBasketByUserID(*userID)
	if err != nil || basket == nil {
		http.Error(w, "Basket not found", http.StatusNotFound)
		return
	}

	wallet, err := db.DB.FetchWalletByUserID(*userID)
	if err != nil || wallet == nil {
		http.Error(w, "Wallet not found", http.StatusNotFound)
		return
	}
	getBasketPrice(basket)
	return
}

func getBasketPrice(basket []model.BasketItem) float32 {
	for _, item := range basket {
		fmt.Println(item)
	}

	return 0
}
