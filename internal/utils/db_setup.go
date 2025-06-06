package utils

import (
	"database/sql"
	"log"
	"os"
)

func SetupTestDB() *sql.DB {
	// Use an in-memory SQLite database for testing
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	queryBytes, err := os.ReadFile("../../internal/database/trivia.sql")
	if err != nil {
		panic(err)
	}
	query := string(queryBytes)
	if _, err := db.Exec(query); err != nil {
		panic(err)
	}

	return db
}
