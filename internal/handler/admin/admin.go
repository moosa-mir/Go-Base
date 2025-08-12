package admin

import "myproject/db"

type Admin struct {
	DB db.AdminDBInterface
}

func NewAdmin(db db.AdminDBInterface) *Admin {
	return &Admin{DB: db}
}
