package route

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"myproject/db"
	"myproject/internal/auth"
	"myproject/internal/handler/account"
	"myproject/internal/handler/admin"
	"myproject/internal/handler/basket"
	"myproject/internal/handler/order"
	"myproject/internal/handler/payment"
	"myproject/internal/handler/product"
	"myproject/internal/handler/seller"
	"myproject/internal/handler/wallet"
	"net/http"
)

// RegisterRoutes registers the login route with the HTTP server
func RegisterRoutes() error {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return err
	}
	defer database.Close()

	fmt.Println("Registering routes...")
	router := mux.NewRouter()
	accountDB := account.NewAccount(database)
	productDB := product.NewProduct(database)
	basketDB := basket.NewBasket(database)
	walletDB := wallet.ApiHandler(database)
	paymentDB := payment.NewPayment(database)
	adminDB := admin.NewAdmin(database)
	sellerDB := seller.NewSeller(database)
	orderDB := order.NewOrder(database)

	// mux.HandleFunc("/login", login.LoginHandler)
	router.HandleFunc("/login", accountDB.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/register", accountDB.RegisterHandler).Methods(http.MethodPost)
	router.HandleFunc("/product", productDB.ProductListHandler).Methods(http.MethodGet)
	router.HandleFunc("/product/{id}", productDB.ProductItemHandler).Methods(http.MethodGet)

	router.HandleFunc("/seller/register", sellerDB.RegisterHandler).Methods(http.MethodPost)

	router.HandleFunc("/admin/login", adminDB.AdminLoginHandler).Methods(http.MethodPost)

	router.HandleFunc("/basket", auth.Middleware(basketDB.WalletItemsHandler)).Methods(http.MethodGet)
	router.HandleFunc("/basket", auth.Middleware(basketDB.AddItemHandler)).Methods(http.MethodPost)
	router.HandleFunc("/basket", auth.Middleware(basketDB.DeleteItemHandle)).Methods(http.MethodDelete)

	router.HandleFunc("/wallet", auth.Middleware(walletDB.HandleIncreaseWallet)).Methods(http.MethodPost)
	router.HandleFunc("/wallet", auth.Middleware(walletDB.WalletHandler)).Methods(http.MethodGet)

	router.HandleFunc("/account", auth.Middleware(accountDB.AccountInfoHandler)).Methods(http.MethodGet)
	router.HandleFunc("/account", auth.Middleware(accountDB.UpdateHandler)).Methods(http.MethodPatch)

	router.HandleFunc("/pay", auth.Middleware(paymentDB.PayHandler)).Methods(http.MethodPost)
	
	router.HandleFunc("/admin/sellers", auth.Middleware(adminDB.SellerListHandler)).Methods(http.MethodGet)
	router.HandleFunc("/admin/unpayouts", auth.Middleware(adminDB.FetchUnPayoutHandler)).Methods(http.MethodGet)

	router.HandleFunc("/order", auth.Middleware(orderDB.ListOrdersHandler)).Methods(http.MethodGet)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
	return nil
}
