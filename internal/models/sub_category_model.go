package models

import (
	"database/sql"
	"time"
)

type SubCategory struct {
	ID         int            `json:"id"`
	CategoryID int            `json:"category_id"`
	Name       string         `json:"game_name"`
	ImageUrl   sql.NullString `json:"score"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  sql.NullTime   `json:"updated_at"`
}
