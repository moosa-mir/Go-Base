package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	model "myproject/internal/model"
)

func (db *DB) FetchAdminByUsername(username string) (*model.StoredAdmin, error) {
	query := `SELECT id, username, password FROM admins WHERE username = $1`
	row := db.QueryRow(query, username)

	var model model.StoredAdmin
	if err := row.Scan(&model.ID, &model.Username, &model.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Admin with username %s not found", username)
		}
		return nil, fmt.Errorf("failed to fetch admin: %w", err)
	}

	log.Printf("Fetched admin: %+v\n", model)
	return &model, nil
}
