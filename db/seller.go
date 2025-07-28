package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	model "myproject/internal/model"
	"myproject/internal/utils"
	"time"
)

func (db *DB) FetchSellerByUsername(username string) (*model.StoredSeller, error) {
	query := `SELECT id, username, password, name, family, phone, account_number, address FROM sellers WHERE username = $1`
	row := db.QueryRow(query, username)

	var seller model.StoredSeller
	if err := row.Scan(&seller.ID, &seller.Username, &seller.Password, &seller.Name, &seller.Family, &seller.Phone, &seller.AccountNumber, &seller.Address); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Seller with username %s not found", username)
		}
		return nil, fmt.Errorf("failed to fetch seller: %w", err)
	}

	log.Printf("Fetched seller: %+v\n", seller)
	return &seller, nil
}

func (db *DB) InsertSeller(seller model.RegistrationSeller) (bool, error) {
	// Hash the password before storing it in the database
	hashedPassword := utils.GetHashPassword(seller.Password)
	if hashedPassword == "" {
		log.Printf("Error hashing password for user %s", seller.Username)
		return false, fmt.Errorf("error hashing password")
	}

	// Define the SQL query with PostgresSQL placeholders ($1, $2, ...)
	insertSQL := `
        INSERT INTO sellers (username, password, name, family, phone, account_number, address)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `

	// Prepare the statement
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	statement, err := db.DB.PrepareContext(ctx, insertSQL)
	if err != nil {
		log.Printf("Error preparing insert statement: %v", err)
		return false, err
	}
	defer statement.Close()

	// Execute the statement with the hashed password and other user details
	_, err = statement.ExecContext(
		ctx,
		seller.Username,
		hashedPassword, // Store the hashed password
		seller.Name,
		seller.Family,
		seller.Phone,
		seller.AccountNumber,
		seller.Address,
	)
	if err != nil {
		log.Printf("Error inserting user for user %s: %v\n", seller.Username, err)
		return false, err
	}

	log.Printf("Inserted user %s\n", seller.Username)
	return true, nil
}

func (db *DB) FetchSellers() ([]model.Seller, error) {
	// Define the SQL query to fetch all products
	query := `
        SELECT id, username, name, family, phone, account_number, address 
        FROM sellers
    `

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error querying products: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store the products
	sellers := []model.Seller{}

	// Iterate over the rows and scan each product into the slice
	for rows.Next() {
		var p model.Seller
		if err := rows.Scan(&p.ID, &p.Username, &p.Name, &p.Family, &p.Phone, &p.AccountNumber, &p.Address); err != nil {
			log.Printf("Error scanning seller row: %v", err)
			return nil, err
		}
		sellers = append(sellers, p)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over product rows: %v", err)
		return nil, err
	}

	// Return the list of products
	return sellers, nil
}
