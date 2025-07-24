package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	product "myproject/internal/model"
	"strings"
)

// FetchProducts fetches a list of products from the database.
func (db *DB) FetchProducts() ([]product.Product, error) {
	// Define the SQL query to fetch all products
	query := `
        SELECT id, name, description, price, image 
        FROM products
    `

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error querying products: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store the products
	products := []product.Product{}

	// Iterate over the rows and scan each product into the slice
	for rows.Next() {
		var p product.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Image); err != nil {
			log.Printf("Error scanning product row: %v", err)
			return nil, err
		}
		products = append(products, p)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over product rows: %v", err)
		return nil, err
	}

	// Return the list of products
	return products, nil
}

// FetchProducts fetches a list of products from the database.
func (db *DB) FetchProduct(productID string) (*product.Product, error) {
	query := `SELECT id, name, description, price, image FROM products WHERE id = $1`
	row := db.QueryRow(query, productID)

	var item product.Product
	if err := row.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Image); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("product with id %v not found", productID)
		}
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	log.Printf("Fetched product: %+v\n", item)
	return &item, nil
}

// FetchProducts fetches a list of products from the database.
func (db *DB) SearchProducts(keyword string) ([]product.Product, error) {
	searchKeyword := "%" + strings.ToLower(keyword) + "%"
	query := `
        SELECT id, name, description, price, image 
        FROM products
        WHERE $1 = '' OR name ILIKE $2
    `
	// Execute the query
	rows, err := db.Query(query, keyword, searchKeyword)
	if err != nil {
		log.Printf("Error querying products: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store the products
	products := []product.Product{}

	// Iterate over the rows and scan each product into the slice
	for rows.Next() {
		var p product.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Image); err != nil {
			log.Printf("Error scanning product row: %v", err)
			return nil, err
		}
		products = append(products, p)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over product rows: %v", err)
		return products, err
	}

	// Return the list of products
	return products, nil
}
