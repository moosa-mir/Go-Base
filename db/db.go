package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"myproject/internal/constant"
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
