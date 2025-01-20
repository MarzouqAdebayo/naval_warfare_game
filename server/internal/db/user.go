package db

import (
	"fmt"
	"github.com/google/uuid"
)

func (db *Database) SaveUser(username string) (uuid.UUID, error) {
	var userID uuid.UUID
	query := `
		INSERT INTO users (username)
		VALUES ($1)
		ON CONFLICT (username)
		DO UPDATE SET updated_at = CURRENT_TIMESTAMP
		RETURNING id;
	`
	err := db.Conn.Raw(query, username).Scan(&userID).Error
	if err != nil {
		return userID, fmt.Errorf("failed to save user: %w", err)
	}
	return userID, nil
}
