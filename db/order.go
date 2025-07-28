package db

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	"myproject/internal/model"
	"time"
)

func (db *DB) MoveItemsToOrders(tx *sql.Tx, userID uuid.UUID, basketItems []model.BasketItem, groupID uuid.UUID) error {
	query := `
        INSERT INTO orders (user_id, product_id, status, group_id, create_date, count)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	createDate := float64(time.Now().Unix())

	for _, item := range basketItems {
		_, err := tx.Exec(query, userID, item.Product.ID, 1, groupID, createDate, item.Count) // Status = 1 (completed)
		if err != nil {
			return fmt.Errorf("failed to insert order: %w", err)
		}
	}

	return nil
}

func (db *DB) FetchOrdersByUserID(userID uuid.UUID) ([]model.OrderItem, error) {
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
        WHERE orders.user_id = $1
    `

	// Execute the query
	rows, err := db.Query(query, userID)
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
