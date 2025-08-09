package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Connect to a database
func ConnectDatabase() *sql.DB {
	var dbPath string

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "local" {
		dbPath = "internal/database/trivia.db"
	} else {
		dbPath = "data/trivia.db"
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	return db
}
