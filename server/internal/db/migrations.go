package db

import (
	"log"
)

// Migrate runs database migrations
func Migrate(db *Database) {
	if err := db.Conn.AutoMigrate(&User{}); err != nil {
		log.Printf("failed to migrate database: %v", err)
		return err
	}
	log.Println("Database migration completed successfully.")
	return nil
}

