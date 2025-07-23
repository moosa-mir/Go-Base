package route

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"myproject/internal/auth"
	"myproject/internal/db"
	"myproject/internal/handler/account"
	"myproject/internal/handler/basket"
	"myproject/internal/handler/product"
	"myproject/internal/handler/wallet"
	"net/http"
)

// RegisterRoutes registers the login route with the HTTP server
func RegisterRoutes(db *db.DB) {
	router := mux.NewRouter()
	accountApi := account.ApiHandler(db)
	productApi := product.ApiHandler(db)
	basketApi := basket.ApiHandler(db)
	walletApi := wallet.ApiHandler(db)

	// mux.HandleFunc("/login", login.LoginHandler)
	router.HandleFunc("/login", accountApi.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/register", accountApi.RegisterHandler).Methods(http.MethodPost)
	router.HandleFunc("/product", productApi.HandleProductList).Methods(http.MethodGet)
	router.HandleFunc("/basket", basketApi.HandleWalletItems).Methods(http.MethodGet)
	router.HandleFunc("/basket", basketApi.HandleAddItem).Methods(http.MethodPost)
	router.HandleFunc("/basket", basketApi.HandleDeleteItem).Methods(http.MethodDelete)
	router.HandleFunc("/wallet", walletApi.HandleIncreaseWallet).Methods(http.MethodPost)
	router.HandleFunc("/wallet", walletApi.HandleWallet).Methods(http.MethodGet)
	router.HandleFunc("/product/{id}", productApi.HandleProductItem).Methods(http.MethodGet)
	router.HandleFunc("/account", auth.Middleware(accountApi.AccountInfoHandler)).Methods(http.MethodGet)
	router.HandleFunc("/account", auth.Middleware(accountApi.UpdateHandler)).Methods(http.MethodPatch)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
