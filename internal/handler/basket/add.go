package basket

import (
	"encoding/json"
	"fmt"
	token "myproject/internal/auth"
	model "myproject/internal/model"
	"myproject/internal/utils"
	"net/http"
)

func (db *Basket) AddItemHandler(w http.ResponseWriter, r *http.Request) {
	userID, errorUserID := token.FetchUserIDFromToken(r)
	if userID == nil || errorUserID != nil {
		http.Error(w, "User name is not valid", http.StatusConflict)
		return
	}
	var insertModel model.InputAddBasketItem
	err := json.NewDecoder(r.Body).Decode(&insertModel)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println("user is adding", insertModel.ProductID, "to db for", userID)
	exist, err := db.DB.IsProductInBasket(insertModel.ProductID, *userID)
	if err != nil || exist == true {
		http.Error(w, "This product already exist on wallet", http.StatusConflict)
		return
	}

	time := utils.GenerateTimeIntervalFromEpoch()
	item := model.InsertBasketItem{ProductID: insertModel.ProductID, Count: 1, Date: time}
	err = db.DB.InsertBasketItem(item, *userID)
	if err != nil {
		http.Error(w, "Error inserting wallet item", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	return
}
