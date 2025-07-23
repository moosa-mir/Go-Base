package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	model "myproject/internal/model"
)

func (db *DB) FetchWalletByUsername(userID uuid.UUID) (*model.WalletModel, error) {
	query := `SELECT amount FROM wallet WHERE username = $1`
	row := db.QueryRow(query, userID)

	var walletModel model.WalletModel
	if err := row.Scan(&walletModel.Amount); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	log.Printf("Fetched wallet: %+v\n", walletModel)
	return &walletModel, nil
}

func (db *DB) IncreaseWallet(item model.IncreaseWalletModel, userID uuid.UUID) error {
	// Define the query to update the wallet amount
	query := `
        INSERT INTO wallet (username, amount)
        VALUES ($1, $2)
        ON CONFLICT (username) 
        DO UPDATE SET amount = wallet.amount + EXCLUDED.amount
    `

	// Execute the query
	_, err := db.Exec(query, userID, item.Amount)
	if err != nil {
		fmt.Println("failed to increase wallet amount: %w", err)
		return fmt.Errorf("failed to increase wallet amount: %w", err)
	}

	fmt.Printf("Increased wallet amount for username %s, %d", userID, item.Amount)
	return nil
}
