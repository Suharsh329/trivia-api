package models

import (
	"database/sql"
	"time"
)

type Game struct {
	ID        int          `json:"id"`
	Name      string       `json:"game_name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
