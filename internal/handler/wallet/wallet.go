package wallet

import (
	"myproject/db"
)

type Api struct {
	DB db.WalletDBInterface
}

func ApiHandler(db db.WalletDBInterface) *Api {
	return &Api{DB: db}
}
