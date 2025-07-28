package product

import (
	"encoding/json"
	"myproject/internal/utils"
	"net/http"
)

func (db *Product) ProductItemHandler(w http.ResponseWriter, r *http.Request) {
	productID, errorPath := utils.GetPathString(r, 2)
	if productID == "" || errorPath != nil {
		http.Error(w, "Product id is not valid", http.StatusNotFound)
		return
	}
	product, error := db.DB.FetchProduct(productID)
	if error != nil {
		http.Error(w, "Product id is not valid", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
