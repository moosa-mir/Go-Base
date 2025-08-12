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
func RegisterRoutes(
	userDB db.UserDBInterface,
	adminDB db.AdminDBInterface,
	basketDB db.BasketDBInterface,
	orderDB db.OrderDBInterface,
	paymentDB db.PaymentDBInterface,
	productDB db.ProductDBInterface,
	sellerDB db.SellerDBInterface,
	walletDB db.WalletDBInterface) error {

	fmt.Println("Registering routes...")
	router := mux.NewRouter()
	newAccount := account.NewAccount(userDB)
	newProduct := product.NewProduct(productDB)
	newBasket := basket.NewBasket(basketDB)
	newWallet := wallet.ApiHandler(walletDB)
	newPayment := payment.NewPayment(paymentDB)
	newAdmin := admin.NewAdmin(adminDB)
	newSeller := seller.NewSeller(sellerDB)
	newOrder := order.NewOrder(orderDB)

	// mux.HandleFunc("/login", login.LoginHandler)
	router.HandleFunc("/login", newAccount.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/register", newAccount.RegisterHandler).Methods(http.MethodPost)
	router.HandleFunc("/product", newProduct.ProductListHandler).Methods(http.MethodGet)
	router.HandleFunc("/product/{id}", newProduct.ProductItemHandler).Methods(http.MethodGet)

	router.HandleFunc("/seller/register", newSeller.RegisterHandler).Methods(http.MethodPost)

	router.HandleFunc("/admin/login", newAdmin.AdminLoginHandler).Methods(http.MethodPost)

	router.HandleFunc("/basket", auth.Middleware(newBasket.WalletItemsHandler)).Methods(http.MethodGet)
	router.HandleFunc("/basket", auth.Middleware(newBasket.AddItemHandler)).Methods(http.MethodPost)
	router.HandleFunc("/basket", auth.Middleware(newBasket.DeleteItemHandle)).Methods(http.MethodDelete)

	router.HandleFunc("/wallet", auth.Middleware(newWallet.HandleIncreaseWallet)).Methods(http.MethodPost)
	router.HandleFunc("/wallet", auth.Middleware(newWallet.WalletHandler)).Methods(http.MethodGet)

	router.HandleFunc("/account", auth.Middleware(newAccount.AccountInfoHandler)).Methods(http.MethodGet)
	router.HandleFunc("/account", auth.Middleware(newAccount.UpdateHandler)).Methods(http.MethodPatch)

	router.HandleFunc("/pay", auth.Middleware(newPayment.PayHandler)).Methods(http.MethodPost)

	router.HandleFunc("/admin/sellers", auth.Middleware(newAdmin.SellerListHandler)).Methods(http.MethodGet)
	router.HandleFunc("/admin/unpayouts", auth.Middleware(newAdmin.FetchUnPayoutHandler)).Methods(http.MethodGet)

	router.HandleFunc("/order", auth.Middleware(newOrder.ListOrdersHandler)).Methods(http.MethodGet)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
	return nil
}
