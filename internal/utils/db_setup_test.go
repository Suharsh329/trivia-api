package utils

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSetupTestDB(t *testing.T) {
	db := SetupTestDB()

	db.Ping()

	db.Close()
}
