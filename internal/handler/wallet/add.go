package wallet

import (
	"encoding/json"
	"fmt"
	token "myproject/internal/auth"
	model "myproject/internal/model"
	"myproject/internal/utils"
	"net/http"
)

func (db *Api) HandleAddItem(w http.ResponseWriter, r *http.Request) {
	username, errorUsername := token.FetchUsernameFromToken(r)
	if username == "" || errorUsername != nil {
		http.Error(w, "User name is not valid", http.StatusConflict)
		return
	}
	var insertModel model.InputAddWalletItem
	err := json.NewDecoder(r.Body).Decode(&insertModel)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println("user is adding", insertModel.ProductID, "to db for", username)
	exist, err := db.DB.IsProductInWallet(insertModel.ProductID, username)
	if err != nil || exist == true {
		http.Error(w, "This product already exist on wallet", http.StatusConflict)
		return
	}

	time := utils.GenerateTimeIntervalFromEpoch()
	item := model.InsertWalletItem{ProductID: insertModel.ProductID, Count: 1, Date: time}
	err = db.DB.InsertWalletItem(item, username)
	if err != nil {
		http.Error(w, "Error inserting wallet item", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	return
}
