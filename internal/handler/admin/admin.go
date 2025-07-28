package admin

import "myproject/db"

type Admin struct {
	DB *db.DB
}

func NewAdmin(db *db.DB) *Admin {
	return &Admin{DB: db}
}
