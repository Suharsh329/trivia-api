package models

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        int            `json:"id"`
	Name      string         `json:"game_name"`
	ImageUrl  sql.NullString `json:"image_url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}
