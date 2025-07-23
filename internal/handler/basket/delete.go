package basket

import (
	"encoding/json"
	"fmt"
	token "myproject/internal/auth"
	"myproject/internal/model"
	"net/http"
)

func (db *Api) HandleDeleteItem(w http.ResponseWriter, r *http.Request) {
	userID, errorUsername := token.FetchUserIDFromToken(r)
	if userID == nil || errorUsername != nil {
		http.Error(w, "User name is not valid", http.StatusConflict)
		return
	}
	var deleteModel model.InputDeleteBasketItem
	err := json.NewDecoder(r.Body).Decode(&deleteModel)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println("user is deleting", deleteModel.ProductID, "on db for", userID)

	err = db.DB.DeleteFromBasket(deleteModel.ProductID, *userID)
	if err != nil {
		http.Error(w, "Error inserting wallet item", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	return
}
