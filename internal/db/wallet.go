package db

import (
	"database/sql"
	"log"
	model "myproject/internal/model"
)

func (db *DB) FetchWalletByUsername(username string) ([]model.WalletItem, error) {
	// FetchProducts fetches a list of products from the database.
	// Define the SQL query to fetch all products
	query := `
        SELECT 
            wallet.id, 
            wallet.count, 
            wallet.date,
            products.id AS product_id,
            products.name AS product_name, 
            products.description AS product_description, 
            products.price AS product_price,
        	products.image AS product_image
        FROM wallet 
        JOIN products ON wallet.product_id = products.id 
        WHERE wallet.username = $1
    `

	// Execute the query
	rows, err := db.Query(query, username)
	if err != nil {
		log.Printf("Error querying wallet: %v", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	items := []model.WalletItem{}
	// Initialize a slice to store the products

	// Iterate over the rows and scan each product into the slice
	for rows.Next() {
		var p model.WalletItem
		if err := rows.Scan(&p.ID, &p.Count, &p.Date, &p.Product.ID, &p.Product.Name, &p.Product.Description, &p.Product.Price, &p.Product.Image); err != nil {
			log.Printf("Error scanning wallet row: %v", err)
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
