package registerRoute

import (
	"fmt"
	"log"
	auth "myproject/Auth"
	db "myproject/DB/Init"
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
	loginApiHandler := login.LoginApiHandler(db)

	// mux.HandleFunc("/login", login.LoginHandler)
	mux.HandleFunc("/login", loginApiHandler.LoginHandler)
	mux.HandleFunc("/register", register.RegisterHandler)
	mux.HandleFunc("/product/", product.HandleProductList)
	mux.HandleFunc("/accountInfo/", auth.AuthMiddleware(account.AccountInfoHandler))
	mux.HandleFunc("/update", auth.AuthMiddleware(update.UpdateHandler))

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
