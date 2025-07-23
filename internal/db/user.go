package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	user "myproject/internal/model"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (db *DB) FetchUserByUserID(userID uuid.UUID) (*user.StoredUser, error) {
	query := `SELECT id, username, name, family, birthday, city, country, phone FROM users WHERE id = $1`
	row := db.QueryRow(query, userID)

	var model user.StoredUser
	if err := row.Scan(&model.ID, &model.Username, &model.Name, &model.Family, &model.Birthday, &model.City, &model.Country, &model.Phone); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with username %s not found", userID)
		}
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	log.Printf("Fetched user: %+v\n", model)
	return &model, nil
}

func (db *DB) FetchUserByUsername(username string) (*user.StoredUser, error) {
	query := `SELECT id, username, name, family, birthday, city, country, phone FROM users WHERE id = $1`
	row := db.QueryRow(query, username)

	var model user.StoredUser
	if err := row.Scan(&model.ID, &model.Username, &model.Name, &model.Family, &model.Birthday, &model.City, &model.Country, &model.Phone); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with username %s not found", username)
		}
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	log.Printf("Fetched user: %+v\n", model)
	return &model, nil
}

// InsertUser inserts a new user into the database.
func (db *DB) InsertUser(user user.RegistrationUser) (bool, error) {
	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password for user %s: %v\n", user.Username, err)
		return false, err
	}

	// Define the SQL query with PostgreSQL placeholders ($1, $2, ...)
	insertSQL := `
        INSERT INTO users (username, password, name, family, birthday, city, country, phone)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
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
		user.Username,
		string(hashedPassword), // Store the hashed password
		user.Name,
		user.Family,
		user.Birthday,
		user.City,
		user.Country,
		user.Phone,
	)
	if err != nil {
		log.Printf("Error inserting user for user %s: %v\n", user.Username, err)
		return false, err
	}

	log.Printf("Inserted user %s\n", user.Username)
	return true, nil
}

func (db *DB) UpdateUser(model user.UpdateUser, userID uuid.UUID) (bool, error) {
	// Define the SQL query with PostgreSQL placeholders ($1, $2, ...)
	updateSQL := `
        UPDATE users 
        SET name = $1, family = $2 
        WHERE username = $3
    `

	// Prepare the statement with context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	statement, err := db.DB.PrepareContext(ctx, updateSQL)
	if err != nil {
		log.Printf("Error preparing update statement: %v", err)
		return false, err
	}
	defer statement.Close()

	// Execute the statement with the provided values
	_, err = statement.ExecContext(ctx, model.Name, model.Family, userID)
	if err != nil {
		log.Printf("Error updating user for user %s: %v\n", userID, err)
		return false, err
	}

	log.Printf("Updated user %s\n", userID)
	return true, nil
}
