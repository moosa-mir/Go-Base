package model

type StoredWalletItem struct {
	ID        string `json:"id"`
	ProductID int    `json:"product_id"`
	Date      string `json:"date"`
	Count     int    `json:"count"`
}

type InsertWalletItem struct {
	ProductID int     `json:"product_id"`
	Date      float32 `json:"date"`
	Count     int     `json:"count"`
}

type InputAddWalletItem struct {
	ProductID int `json:"product_id"`
}

type WalletItem struct {
	ID      string  `json:"id"`
	Product Product `json:"product"`
	Date    string  `json:"date"`
	Count   int     `json:"count"`
}
