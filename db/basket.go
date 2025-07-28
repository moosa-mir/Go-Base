package db

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	model "myproject/internal/model"
)

func (db *DB) FetchBasketByUserID(userID uuid.UUID) ([]model.BasketItem, error) {
	// FetchProducts fetches a list of products from the database.
	// Define the SQL query to fetch all products
	query := `
        SELECT 
            basket.id, 
            basket.count, 
            basket.date,
            products.id AS product_id,
            products.name AS product_name, 
            products.description AS product_description, 
            products.price AS product_price,
        	products.image AS product_image
        FROM basket 
        JOIN products ON basket.product_id = products.id 
        WHERE basket.user_id = $1
    `

	// Execute the query
	rows, err := db.Query(query, userID)
	if err != nil {
		log.Printf("Error querying basket: %v", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	items := []model.BasketItem{}
	// Initialize a slice to store the products

	// Iterate over the rows and scan each product into the slice
	for rows.Next() {
		var p model.BasketItem
		if err := rows.Scan(&p.ID, &p.Count, &p.Date, &p.Product.ID, &p.Product.Name, &p.Product.Description, &p.Product.Price, &p.Product.Image); err != nil {
			log.Printf("Error scanning basket row: %v", err)
			return nil, err
		}
		items = append(items, p)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over product rows: %v", err)
		return nil, err
	}

	// Return the list of products
	return items, nil
}

func (db *DB) IsProductInBasket(productID uuid.UUID, userID uuid.UUID) (bool, error) {
	var exists bool
	query := `
        SELECT EXISTS (
            SELECT 1 
            FROM basket 
            WHERE product_id = $1 AND user_id = $2
        )
    `
	err := db.QueryRow(query, productID, userID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking product in basket: %w", err)
	}
	return exists, nil
}

func (db *DB) InsertBasketItem(item model.InsertBasketItem, userID uuid.UUID) error {
	query := `
        INSERT INTO basket (product_id, count, user_id, date) 
        VALUES ($1, $2, $3, $4)
    `
	_, err := db.Exec(query, item.ProductID, item.Count, userID, item.Date)
	if err != nil {
		return fmt.Errorf("error inserting basket item: %w", err)
	}
	return nil
}

func (db *DB) DeleteFromBasket(productID uuid.UUID, userID uuid.UUID) error {
	// Define the DELETE query
	query := `
        DELETE FROM basket 
        WHERE product_id = $1 AND user_id = $2
    `

	// Execute the query
	_, err := db.Exec(query, productID, userID)
	if err != nil {
		return fmt.Errorf("error deleting item from basket: %w", err)
	}

	return nil
}

func (db *DB) RemoveAllBasketRows(tx *sql.Tx, userID uuid.UUID) error {
	query := `
        DELETE FROM basket 
        WHERE user_id = $1
    `
	_, err := tx.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed to remove basket rows: %w", err)
	}
	return nil
}
