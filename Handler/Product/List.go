package product

import (
	"encoding/json"
	db "myproject/DB"
	"net/http"
)

type ProductApi struct {
	DB *db.DB
}

func ApiHandler(db *db.DB) *ProductApi {
	return &ProductApi{DB: db}
}

func (db *ProductApi) HandleProductList(w http.ResponseWriter, r *http.Request) {
	products, error := db.DB.FetchProducts()
	if error != nil {
		http.Error(w, "Internal server error", http.StatusConflict)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
