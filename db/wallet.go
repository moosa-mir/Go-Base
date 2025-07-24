package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	model "myproject/internal/model"
)

func (db *DB) FetchWalletByUserID(userID uuid.UUID) (*model.WalletModel, error) {
	query := `SELECT balance FROM wallet WHERE user_id = $1`
	row := db.QueryRow(query, userID)

	var walletModel model.WalletModel
	if err := row.Scan(&walletModel.Balance); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	log.Printf("Fetched wallet: %+v\n", walletModel)
	return &walletModel, nil
}

// IncreaseWallet inserts or updates the balance for a specific user.
func (db *DB) IncreaseWallet(balance float32, userID uuid.UUID) error {
	// Define the query to insert or update the balance
	query := `
        INSERT INTO wallet (user_id, balance)
        VALUES ($1, $2)
        ON CONFLICT (user_id) 
        DO UPDATE SET balance = wallet.balance + EXCLUDED.balance
    `

	// Execute the query
	_, err := db.Exec(query, userID, balance)
	if err != nil {
		return fmt.Errorf("failed to increase wallet balance: %w", err)
	}

	fmt.Printf("Increased wallet balance for user_id %s by %.2f\n", userID, balance)
	return nil
}

func (db *DB) UpdateWalletBalance(tx *sql.Tx, userID uuid.UUID, newBalance float32) error {
	query := `
        UPDATE wallet 
        SET balance = balance + $1 
        WHERE user_id = $2
    `
	result, err := tx.Exec(query, newBalance, userID)
	if err != nil {
		return fmt.Errorf("failed to update wallet balance: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("wallet not found for user_id: %s", userID)
	}

	return nil
}
