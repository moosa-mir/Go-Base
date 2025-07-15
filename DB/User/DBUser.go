package db

import (
	_ "github.com/mattn/go-sqlite3"
)

// Init initializes the database connection and creates the tokens table if not exists
func Init() {
	// db.Init()
	// CreateUserTable()
}

// CreateTokenTable creates the tokens table if it doesn't exist
// func CreateUserTable() {
// 	createTableSQL := `
//     CREATE TABLE IF NOT EXISTS users (
//         user_id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		username TEXT NOT NULL UNIQUE,
// 		password TEXT NOT NULL,
//         name TEXT NOT NULL,
// 		family TEXT NOT NULL,
// 		birthday TEXT NOT NULL,
// 		city TEXT NOT NULL,
//         country TEXT NOT NULL
//     );`

// 	statement, err := db.DB.Prepare(createTableSQL)
// 	if err != nil {
// 		log.Fatalf("Error preparing statement: %v", err)
// 	}
// 	_, err = statement.Exec()
// 	if err != nil {
// 		log.Fatalf("Error creating users table: %v", err)
// 	}
// 	log.Println("Users table created or already exists")
// }

// InsertTokenForUserID inserts a new token for a user
