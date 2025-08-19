package payout

import (
	"fmt"
	"myproject/db"
)

func PayOutHandler(dbProvider *db.DatabaseProvider) {
	unPayoutOrders, err := dbProvider.AdminDB.FetchUnPayoutOrders()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = dbProvider.PayoutDB.Payout(unPayoutOrders)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Payout successfuly done ")
}
