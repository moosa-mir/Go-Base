package db

import (
	"database/sql"
	"fmt"
	"log"
	"myproject/internal/model"
)

func (db *DB) FetchUnPayoutOrders() ([]model.OrderItem, error) {
	// FetchProducts fetches a list of products from the database.
	// Define the SQL query to fetch all products
	query := `
        SELECT 
            orders.id,
            orders.status,
            orders.group_id,
            orders.create_date,
            orders.count,
            products.id AS product_id,
            products.name AS product_name, 
            products.description AS product_description, 
            products.price AS product_price,
        	products.image AS product_image
        FROM orders 
        JOIN products ON orders.product_id = products.id 
        WHERE orders.seller_status = 0
    `

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error querying order: %v", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	items := []model.OrderItem{}
	// Initialize a slice to store the products

	// Iterate over the rows and scan each product into the slice
	for rows.Next() {
		var p model.OrderItem
		if err := rows.Scan(&p.ID, &p.Status, &p.GroupID, &p.CreateDate, &p.Count, &p.Product.ID, &p.Product.Name, &p.Product.Description, &p.Product.Price, &p.Product.Image); err != nil {
			log.Printf("Error scanning order row: %v", err)
			return nil, err
		}
		items = append(items, p)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over orders rows: %v", err)
		return nil, err
	}

	// Return the list of products
	return items, nil
}

// payout updates the seller_status to 1 for all items in the provided OrderItem slice
func (db *DB) Payout(orderItems []model.OrderItem) error {
	// Prepare the SQL query to update seller_status
	query := "UPDATE orders SET seller_status = 1 WHERE id = $1"

	// Start a transaction to ensure atomicity
	tx, err := db.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Rollback in case of error
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			log.Printf("Transaction rolled back due to error: %v", err)
		}
	}()

	// Iterate over the OrderItem slice and update seller_status for each item
	for _, item := range orderItems {
		_, err := tx.Exec(query, item.ID)
		if err != nil {
			return fmt.Errorf("failed to update seller_status for item ID %d: %w", item.ID, err)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Seller status updated successfully for all items")
	return nil
}
