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
	totalPrice := getBasketPrice(basket)
	fmt.Println("totalPrice of basket is", totalPrice)
	if wallet.Balance < totalPrice {
		http.Error(w, "Not enough balance", http.StatusBadRequest)
		return
	}

	errorMoveToDB := db.DB.ProcessPayment(*userID, basket, totalPrice)
	if errorMoveToDB != nil {
		http.Error(w, errorMoveToDB.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func getBasketPrice(basket []model.BasketItem) float32 {
	var totalPrice float32 = 0
	for _, item := range basket {
		totalPrice += item.Product.Price * float32(item.Count)
	}

	return totalPrice
}
