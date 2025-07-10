package accountinfo

import (
	json "encoding/json"
	"fmt"
	dbUser "myproject/DB/User"
	token "myproject/Token"
	http "net/http"
	"strconv"
	"strings"
)

func AccountInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is Get
	fmt.Println("AccountInfoHandler")
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Extract the path parameter "id"
	accountID, error := getAccountID(r)
	if error != nil {
		fmt.Println("No account ID found")
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	isTokenValid, error := token.CheckToken(r, accountID)
	if !isTokenValid {
		fmt.Println(error)
		http.Error(w, "User invalid", http.StatusUnauthorized)
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

func getAccountID(r *http.Request) (int, error) {
	// Extract the path parameter manually
	path := r.URL.Path
	fmt.Println("path is ", path)
	parts := strings.Split(path, "/") // Split the path into parts
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid URL")
	}

	accountIDStr := parts[2] // Extract the "id" from the path
	// Convert the account ID to an integer
	accountID, err := strconv.Atoi(accountIDStr)
	fmt.Println("string account id is ", accountIDStr)
	if err != nil {
		return 0, fmt.Errorf("invalid URL")
	}
	return accountID, nil
}
