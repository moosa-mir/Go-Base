package route

import (
	"fmt"
	"github.com/gorilla/mux"
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
	router := mux.NewRouter()
	accountApi := account.ApiHandler(db)
	productApi := product.ApiHandler(db)
	walletApi := wallet.ApiHandler(db)

	// mux.HandleFunc("/login", login.LoginHandler)
	router.HandleFunc("/login", accountApi.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/register", accountApi.RegisterHandler).Methods(http.MethodPost)
	router.HandleFunc("/product", productApi.HandleProductList).Methods(http.MethodGet)
	router.HandleFunc("/wallet", walletApi.HandleWalletItems).Methods(http.MethodGet)
	router.HandleFunc("/wallet", walletApi.HandleAddItem).Methods(http.MethodPost)
	router.HandleFunc("/product/{id}", productApi.HandleProductItem).Methods(http.MethodGet)
	router.HandleFunc("/accountInfo/", auth.AuthMiddleware(accountApi.AccountInfoHandler)).Methods(http.MethodGet)
	router.HandleFunc("/update", auth.AuthMiddleware(accountApi.UpdateHandler)).Methods(http.MethodPatch)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
