package route

import (
	"fmt"
	"log"
	"myproject/internal/auth"
	"myproject/internal/db"
	"myproject/internal/handler/account"
	"myproject/internal/handler/product"
	"myproject/internal/handler/wallet"
	"net/http"
)

// RegisterRoutes registers the login route with the HTTP server
func RegisterRoutes(db *db.DB) {
	mux := http.NewServeMux()
	accountApi := account.ApiHandler(db)
	productApi := product.ApiHandler(db)
	walletApi := wallet.ApiHandler(db)

	// mux.HandleFunc("/login", login.LoginHandler)
	mux.HandleFunc("/login", accountApi.LoginHandler)
	mux.HandleFunc("/register", accountApi.RegisterHandler)
	mux.HandleFunc("/product", productApi.HandleProductList)
	mux.HandleFunc("/wallet", walletApi.HandleWalletItems)
	mux.HandleFunc("/product/{id}", productApi.HandleProductItem)
	mux.HandleFunc("/accountInfo/", auth.AuthMiddleware(accountApi.AccountInfoHandler))
	mux.HandleFunc("/update", auth.AuthMiddleware(accountApi.UpdateHandler))

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
