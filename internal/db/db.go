package db

import (
	"database/sql"
	"fmt"
	"log"
	"myproject/internal/constant"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type DB struct {
	*sql.DB
}

func ConnectDB() (*DB, error) {
	db, err := sql.Open("postgres", constant.ConnStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to the database!")
	return &DB{db}, nil
}

func (db *DB) Close() {
	db.DB.Close()
}
