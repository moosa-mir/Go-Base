package model

type StoredWalletItem struct {
	ID        string `json:"id"`
	ProductID int    `json:"product_id"`
	Date      string `json:"date"`
	Count     int    `json:"count"`
}

type InsertWalletItem struct {
	ProductID int    `json:"product_id"`
	Username  int    `json:"username"`
	Date      string `json:"date"`
	Count     int    `json:"count"`
}

type WalletItem struct {
	ID      string  `json:"id"`
	Product Product `json:"product"`
	Date    string  `json:"date"`
	Count   int     `json:"count"`
}
