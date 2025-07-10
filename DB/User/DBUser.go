package user

import (
	sql "database/sql"
	"fmt"
	"log"
	db "myproject/DB"
	user "myproject/User"

	_ "github.com/mattn/go-sqlite3"
)

// Init initializes the database connection and creates the tokens table if not exists
func Init() {
	db.Init()
	CreateUserTable()
}

// CreateTokenTable creates the tokens table if it doesn't exist
func CreateUserTable() {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
        name TEXT NOT NULL,
		family TEXT NOT NULL,
		birthday TEXT NOT NULL,
		city TEXT NOT NULL,
        country TEXT NOT NULL
    );`

	statement, err := db.DB.Prepare(createTableSQL)
	if err != nil {
		log.Fatalf("Error preparing statement: %v", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}
	log.Println("Users table created or already exists")
}

// InsertTokenForUserID inserts a new token for a user
func InsertUser(user user.RegistrationUser) {
	insertSQL := `INSERT INTO users (username, password, name, family, birthday, city, country) VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.DB.Prepare(insertSQL)
	if err != nil {
		log.Fatalf("Error preparing insert statement: %v", err)
	}
	defer statement.Close()

	_, err = statement.Exec(user.Username, user.Password, user.Name, user.Family, user.Birthday, user.City, user.Country)
	if err != nil {
		log.Printf("Error inserting user for user %s: %v\n", user.Username, err)
		return
	}
	log.Printf("Inserted user %s\n", user.Username)
}

// FetchTokenByUserID retrieves the token associated with a user_id
func FetchUserByUserID(userID int) (*user.User, error) {
	var user user.User
	err := db.DB.QueryRow("SELECT user_id, username, name, family, Birthday, city, country FROM users WHERE user_id = ?", userID).Scan(&user.UserID, &user.Username, &user.Name, &user.Family, &user.Birthday, &user.City, &user.Country)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no token found for user_id %d", userID)
		}
		return nil, fmt.Errorf("database error: %w", err)
	}
	return &user, nil
}

// FetchTokenByUserID retrieves the token associated with a user_id
func FetchUserByUsername(username string) (*user.User, error) {
	var storedUser user.StoredUser
	err := db.DB.QueryRow("SELECT token FROM users WHERE username = ?", username).Scan(&storedUser)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no token found for username %s", username)
		}
		return nil, fmt.Errorf("database error: %w", err)
	}
	userObject := storedUser.ConvertToUser()
	return &userObject, nil
}
