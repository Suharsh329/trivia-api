package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Connect to a database
func ConnectDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "internal/database/trivia.db")
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	return db
}
