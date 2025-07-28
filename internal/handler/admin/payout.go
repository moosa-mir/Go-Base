package admin

import "net/http"

func (db *Admin) PayoutHandler(w http.ResponseWriter, r *http.Request) {
	db.DB.Payout()
}
