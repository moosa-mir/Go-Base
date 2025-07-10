package token

import (
	sql "database/sql"
	"fmt"
	"log"
	db "myproject/DB"

	_ "github.com/mattn/go-sqlite3"
)

// Init initializes the database connection and creates the tokens table if not exists
func Init() {
	db.Init()
	CreateTokenTable()
}

// CreateTokenTable creates the tokens table if it doesn't exist
func CreateTokenTable() {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS tokens (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        token TEXT NOT NULL UNIQUE
    );`

	statement, err := db.DB.Prepare(createTableSQL)
	if err != nil {
		log.Fatalf("Error preparing statement: %v", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalf("Error creating tokens table: %v", err)
	}
	log.Println("Tokens table created or already exists")
}

// InsertTokenForUserID inserts a new token for a user
func InsertTokenForUserID(userID int, token string) {
	insertSQL := `INSERT INTO tokens (user_id, token) VALUES (?, ?)`
	statement, err := db.DB.Prepare(insertSQL)
	if err != nil {
		log.Fatalf("Error preparing insert statement: %v", err)
	}
	defer statement.Close()

	_, err = statement.Exec(userID, token)
	if err != nil {
		log.Printf("Error inserting token for user %d: %v\n", userID, err)
		return
	}
	log.Printf("Inserted token for user %d\n", userID)
}

// FetchTokenByUserID retrieves the token associated with a user_id
func FetchTokenByUserID(userID int) (string, error) {
	var token string
	err := db.DB.QueryRow("SELECT token FROM tokens WHERE user_id = ?", userID).Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no token found for user_id %d", userID)
		}
		return "", fmt.Errorf("database error: %w", err)
	}

	return token, nil
}

// FetchUserIDOfToken gets the user_id associated with the given token
func FetchUserIDOfToken(token string) (int, error) {
	var userID int
	err := db.DB.QueryRow("SELECT user_id FROM tokens WHERE token = ?", token).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no user found for token")
		}
		return 0, fmt.Errorf("database error: %w", err)
	}

	return userID, nil
}

// IsTokenValid checks if a token exists
func IsTokenValid(token string) bool {
	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM tokens WHERE token = ?)", token).Scan(&exists)
	if err != nil {
		log.Printf("Error validating token: %v\n", err)
		return false
	}
	return exists
}
