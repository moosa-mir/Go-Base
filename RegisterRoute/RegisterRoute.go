package registerRoute

import (
	"fmt"
	"log"
	auth "myproject/Auth"
	db "myproject/DB"
	account "myproject/Handler/Account"
	product "myproject/Handler/Product"
	"net/http"
)

// RegisterRoutes registers the login route with the HTTP server
func RegisterRoutes(db *db.DB) {
	mux := http.NewServeMux()
	accountApi := account.ApiHandler(db)
	productApi := product.ApiHandler(db)

	// mux.HandleFunc("/login", login.LoginHandler)
	mux.HandleFunc("/login", accountApi.LoginHandler)
	mux.HandleFunc("/register", accountApi.RegisterHandler)
	mux.HandleFunc("/product", productApi.HandleProductList)
	mux.HandleFunc("/product/{id}", productApi.HandleProductItem)
	mux.HandleFunc("/accountInfo/", auth.AuthMiddleware(accountApi.AccountInfoHandler))
	mux.HandleFunc("/update", auth.AuthMiddleware(accountApi.UpdateHandler))

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
