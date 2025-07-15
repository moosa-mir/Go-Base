package registerRoute

import (
	"fmt"
	"log"
	auth "myproject/Auth"
	db "myproject/DB"
	product "myproject/Handler/Product/List"
	account "myproject/Handler/User/AccountInfo"
	login "myproject/Handler/User/Login"
	register "myproject/Handler/User/Register"
	update "myproject/Handler/User/Update"
	"net/http"
)

// RegisterRoutes registers the login route with the HTTP server
func RegisterRoutes(db *db.DB) {
	mux := http.NewServeMux()
	loginApi := login.ApiHandler(db)
	registerApi := register.ApiHandler(db)
	accountApi := account.ApiHandler(db)
	updateApi := update.ApiHandler(db)

	// mux.HandleFunc("/login", login.LoginHandler)
	mux.HandleFunc("/login", loginApi.LoginHandler)
	mux.HandleFunc("/register", registerApi.RegisterHandler)
	mux.HandleFunc("/product/", product.HandleProductList)
	mux.HandleFunc("/accountInfo/", auth.AuthMiddleware(accountApi.AccountInfoHandler))
	mux.HandleFunc("/update", auth.AuthMiddleware(updateApi.UpdateHandler))

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
