package db

import (
	sql "database/sql"
	"log"
	constant "myproject/Constant"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Init initializes the database connection and creates the tokens table if not exists
func Init() {
	var err error
	DB, err = sql.Open("sqlite3", constant.DataSource)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Connected to SQLite database")
}
