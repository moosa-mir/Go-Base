package db

import (
	"database/sql"
	"github.com/google/uuid"
	"myproject/internal/model"
)

type UserDBInterface interface {
	FetchUserByUserID(userID uuid.UUID) (*model.StoredUser, error)
	FetchUserByUsername(username string) (*model.StoredUser, error)
	InsertUser(user model.RegistrationUser) (bool, error)
	UpdateUser(model model.UpdateUser, userID uuid.UUID) (bool, error)
}

type SellerDBInterface interface {
	FetchSellerByUsername(username string) (*model.StoredSeller, error)
	InsertSeller(seller model.RegistrationSeller) (bool, error)
	FetchSellers() ([]model.Seller, error)
}

type WalletDBInterface interface {
	IncreaseWallet(balance float32, userID uuid.UUID) error
	UpdateWalletBalance(tx *sql.Tx, userID uuid.UUID, newBalance float32) error
	FetchWalletByUserID(userID uuid.UUID) (*model.WalletModel, error)
}

type ProductDBInterface interface {
	FetchProducts() ([]model.Product, error)
	FetchProduct(productID string) (*model.Product, error)
	SearchProducts(keyword string) ([]model.Product, error)
}

type Payout interface {
	Payout(orderItems []model.OrderItem) error
}

type PaymentDBInterface interface {
	ProcessPayment(userID uuid.UUID, basketItems []model.BasketItem, totalPrice float32) error
	FetchBasketByUserID(userID uuid.UUID) ([]model.BasketItem, error)
	FetchWalletByUserID(userID uuid.UUID) (*model.WalletModel, error)
}

type OrderDBInterface interface {
	MoveItemsToOrders(tx *sql.Tx, userID uuid.UUID, basketItems []model.BasketItem, groupID uuid.UUID) error
	FetchOrdersByUserID(userID uuid.UUID) ([]model.OrderItem, error)
}

type BasketDBInterface interface {
	IsProductInBasket(productID uuid.UUID, userID uuid.UUID) (bool, error)
	InsertBasketItem(item model.InsertBasketItem, userID uuid.UUID) error
	DeleteFromBasket(productID uuid.UUID, userID uuid.UUID) error
	RemoveAllBasketRows(tx *sql.Tx, userID uuid.UUID) error
	FetchBasketByUserID(userID uuid.UUID) ([]model.BasketItem, error)
}

type AdminDBInterface interface {
	FetchAdminByUsername(username string) (*model.StoredAdmin, error)
	FetchUnPayoutOrders() ([]model.OrderItem, error)
	FetchSellers() ([]model.Seller, error)
}
