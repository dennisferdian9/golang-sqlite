package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database
func InitDB() {
	// Get the database file from the environment variables
	dbFile := os.Getenv("DB_FILE")

	if dbFile == "" {
		log.Fatalf("DB_FILE environment variable not set")
	}

	var err error
	DB, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	createTable()
}

// Create the users table
func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		username TEXT PRIMARY KEY,
		name TEXT
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
