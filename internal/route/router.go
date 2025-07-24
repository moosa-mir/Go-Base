package route

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"myproject/db"
	"myproject/internal/auth"
	"myproject/internal/handler/account"
	"myproject/internal/handler/basket"
	"myproject/internal/handler/payment"
	"myproject/internal/handler/product"
	"myproject/internal/handler/wallet"
	"net/http"
)

// RegisterRoutes registers the login route with the HTTP server
func RegisterRoutes(db *db.DB) {
	router := mux.NewRouter()
	accountApi := account.NewAccount(db)
	productApi := product.ApiHandler(db)
	basketApi := basket.NewBasket(db)
	walletApi := wallet.ApiHandler(db)
	paymentApi := payment.NewPayment(db)

	// mux.HandleFunc("/login", login.LoginHandler)
	router.HandleFunc("/login", accountApi.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/register", accountApi.RegisterHandler).Methods(http.MethodPost)
	router.HandleFunc("/product", productApi.ProductListHandler).Methods(http.MethodGet)
	router.HandleFunc("/product/{id}", productApi.ProductItemHandler).Methods(http.MethodGet)

	router.HandleFunc("/basket", auth.Middleware(basketApi.WalletItemsHandler)).Methods(http.MethodGet)
	router.HandleFunc("/basket", auth.Middleware(basketApi.AddItemHandler)).Methods(http.MethodPost)
	router.HandleFunc("/basket", auth.Middleware(basketApi.DeleteItemHandle)).Methods(http.MethodDelete)

	router.HandleFunc("/wallet", auth.Middleware(walletApi.HandleIncreaseWallet)).Methods(http.MethodPost)
	router.HandleFunc("/wallet", auth.Middleware(walletApi.WalletHandler)).Methods(http.MethodGet)

	router.HandleFunc("/account", auth.Middleware(accountApi.AccountInfoHandler)).Methods(http.MethodGet)
	router.HandleFunc("/account", auth.Middleware(accountApi.UpdateHandler)).Methods(http.MethodPatch)

	router.HandleFunc("/pay", auth.Middleware(paymentApi.PayHandler)).Methods(http.MethodPost)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
