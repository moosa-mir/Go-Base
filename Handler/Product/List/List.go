package list

import (
	utils "myproject/Utils"
	"net/http"
)

func HandleProductList(w http.ResponseWriter, r *http.Request) {
	productID, error := utils.GetPathInt(r, 2)
	if error != nil || productID == 0 {
		http.Error(w, "Product ID is not valid", http.StatusBadRequest)
		return
	}
}
