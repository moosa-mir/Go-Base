package basket

import (
	"myproject/internal/db"
)

type Api struct {
	DB *db.DB
}

func ApiHandler(db *db.DB) *Api {
	return &Api{DB: db}
}
