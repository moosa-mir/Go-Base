package db

import (
	"fmt"
	"github.com/google/uuid"
	"myproject/internal/model"
)

func (db *DB) ProcessPayment(userID uuid.UUID, basketItems []model.BasketItem, totalPrice float32) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	// Rollback the transaction if an error occurs
	defer func() {
		if err != nil {
			tx.Rollback()
			fmt.Println("Transaction rolled back due to error:", err)
		}
	}()

	// Generate a unique group ID for the orders
	groupID := uuid.New()

	// 1. Move items to orders
	err = db.MoveItemsToOrders(tx, userID, basketItems, groupID)
	if err != nil {
		return fmt.Errorf("failed to move items to orders: %w", err)
	}

	// 2. Remove all basket rows
	err = db.RemoveAllBasketRows(tx, userID)
	if err != nil {
		return fmt.Errorf("failed to remove basket rows: %w", err)
	}

	// 3. Update wallet balance
	err = db.UpdateWalletBalance(tx, userID, -totalPrice) // Deduct the total amount from the wallet
	if err != nil {
		return fmt.Errorf("failed to update wallet balance: %w", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Payment processed successfully")
	return nil
}
