package product

import (
	"encoding/json"
	"myproject/internal/utils"
	"net/http"
)

func (db *Product) ProductListHandler(w http.ResponseWriter, r *http.Request) {
	keyword := utils.GetQueryParam(r, "keyword")
	if keyword == "" {
		products, error := db.DB.FetchProducts()
		if error != nil {
			http.Error(w, "Internal server error", http.StatusConflict)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	} else {
		products, error := db.DB.SearchProducts(keyword)
		if error != nil {
			http.Error(w, "Internal server error", http.StatusConflict)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}
