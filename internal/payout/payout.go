package payout

import (
	"fmt"
	"myproject/db"
)

func PayOutHandler(db *db.DB) {
	unPayoutOrders, err := db.FetchUnPayoutOrders()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Payout(unPayoutOrders)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Payout successfuly done ")
}
