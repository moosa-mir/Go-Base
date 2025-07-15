package accountinfo

import (
	json "encoding/json"
	"fmt"
	dbUser "myproject/DB/User"
	utils "myproject/Utils"
	http "net/http"
)

func AccountInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is Get
	fmt.Println("AccountInfoHandler")
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Extract the path parameter "id"
	accountID, error := utils.GetPathInt(r, 2)
	if error != nil {
		fmt.Println("No account ID found")
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	user, error := dbUser.FetchUserByUserID(accountID)
	if error != nil || user == nil {
		fmt.Println(error)
		http.Error(w, "Invalid Account", http.StatusBadRequest)
		return
	}

	fmt.Println("account ID is", accountID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
