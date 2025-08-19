package db

// DatabaseProvider holds all database interfaces separately
type DatabaseProvider struct {
	UserDB     UserDBInterface
	AdminDB    AdminDBInterface
	BasketDB   BasketDBInterface
	OrderDB    OrderDBInterface
	PaymentDB  PaymentDBInterface
	ProductDB  ProductDBInterface
	SellerDB   SellerDBInterface
	WalletDB   WalletDBInterface
	PayoutDB   Payout
}

// NewDatabaseProvider creates a new database provider with all interfaces
func NewDatabaseProvider(db *DB) *DatabaseProvider {
	return &DatabaseProvider{
		UserDB:     db,
		AdminDB:    db,
		BasketDB:   db,
		OrderDB:    db,
		PaymentDB:  db,
		ProductDB:  db,
		SellerDB:   db,
		WalletDB:   db,
		PayoutDB:   db,
	}
}

