package models

import (
	"database/sql"
	"time"
)

type Team struct {
	ID        int          `json:"id"`
	GameID    int          `json:"game_id"`
	Name      string       `json:"game_name"`
	Score     float64      `json:"score"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
