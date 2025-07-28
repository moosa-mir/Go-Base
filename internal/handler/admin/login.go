package admin

import (
	"encoding/json"
	"fmt"
	token "myproject/internal/auth"
	"myproject/internal/handler/account"
	"myproject/internal/model"
	"myproject/internal/utils"
	"net/http"
)

func (db *Admin) AdminLoginHandler(w http.ResponseWriter, r *http.Request) {

	var loginModel model.LoginRequest
	errDecode := json.NewDecoder(r.Body).Decode(&loginModel)
	if errDecode != nil {
		http.Error(w, errDecode.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Login request is ", loginModel)
	adminOnDBModel, err := db.DB.FetchAdminByUsername(loginModel.Username)
	if err != nil || adminOnDBModel == nil {
		http.Error(w, fmt.Sprintf("failed to fetch admin by username: %s", loginModel.Username), http.StatusBadRequest)
		return
	}

	hashedPassword := utils.GetHashPassword(loginModel.Password)
	fmt.Println("Hash password is", hashedPassword)
	if hashedPassword == "" {
		http.Error(w, fmt.Sprintf("failed to hash password: %s", loginModel.Password), http.StatusBadRequest)
		return
	}

	if adminOnDBModel.Password != hashedPassword {
		http.Error(w, "username does not match", http.StatusBadRequest)
		return
	}

	newToken, err := token.GenerateToken(adminOnDBModel.Username, adminOnDBModel.ID)
	if err != nil || newToken == "" {
		http.Error(w, fmt.Sprintf("failed to generate token: %s", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := account.TokenResponse{
		Token: newToken,
	}
	_ = json.NewEncoder(w).Encode(response)
	return
}
