package models

import (
	"database/sql"
	"time"
)

type Question struct {
	ID               int            `json:"id"`
	SubCategoryID    int            `json:"sub_category_id"`
	QuestionText     string         `json:"question_text"`
	CorrectAnswer    string         `json:"correct_answer"`
	AcceptableAnswer sql.NullString `json:"acceptable_answer"`
	DifficultyLevel  int            `json:"difficulty_level"`
	ImageURL         sql.NullString `json:"image_url"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        sql.NullTime   `json:"updated_at"`
}
