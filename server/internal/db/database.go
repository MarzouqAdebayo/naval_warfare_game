package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

// var DB *gorm.DB
type Database struct {
	Conn *gorm.DB
}

// Init initializes the database connection
func Init() (*Database, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Error loading .env file %v", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Error("DATABASE_URL environment variable is not set")
	}

	// Open a connection to the database
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	log.Println("Database connected successfully.")

	return &Database{Conn: conn}, nil
}
