package models

import (
	"database/sql"
	"time"
)

type QuestionQueue struct {
	ID         int          `json:"id"`
	GameID     int          `json:"game_id"`
	QuestionID int          `json:"question_id"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}

type QuestionQueueResponse struct {
	QueueID          int     `json:"queue_id"`
	QuestionID       int     `json:"question_id"`
	QuestionText     string  `json:"question_text"`
	CorrectAnswer    string  `json:"correct_answer"`
	AcceptableAnswer string  `json:"acceptable_answer"`
	DifficultyLevel  int     `json:"difficulty_level"`
	ImageURL         *string `json:"image_url"`
	SubCategoryID    int     `json:"sub_category_id"`
	SubCategoryName  string  `json:"sub_category_name"`
	GameID           int     `json:"game_id"`
}
