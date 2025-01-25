package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"trivia/internal/models"
	"trivia/internal/utils"
)

type QuestionRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) Get(filters map[string]any) ([]models.Question, error) {
	sql := `SELECT * FROM questions`

	if len(filters) > 0 {
		exactMatch, ok := filters["exactMatch"].(bool)
		if !ok {
			exactMatch = false
		}
		sql += utils.CreateQueryFilters(filters, exactMatch)
	}

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return nil, fmt.Errorf("QuestionRepository@Get | Error: %v", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("QuestionRepository@Get | Error: %v", err)
	}
	defer rows.Close()

	questions := []models.Question{}
	for rows.Next() {
		question := models.Question{}
		if err := rows.Scan(&question.ID, &question.SubCategoryID, &question.QuestionText, &question.CorrectAnswer, &question.AcceptableAnswer, &question.DifficultyLevel, &question.ImageURL, &question.CreatedAt, &question.UpdatedAt); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("QuestionRepository@Get | Error: %v", err)
	}

	return questions, nil
}

func (r *QuestionRepository) Create(question models.Question) (int, error) {
	sql := `INSERT INTO questions(sub_category_id, question_text, correct_answer, acceptable_answer, difficulty_level, image_url) VALUES(?, ?, ?, ?, ?, ?) RETURNING id`

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return 0, fmt.Errorf("QuestionRepository@Create | Error: %v", err)
	}

	var id int

	if err := stmt.QueryRow(question.SubCategoryID, question.QuestionText, question.CorrectAnswer, question.AcceptableAnswer, question.DifficultyLevel, question.ImageURL).Scan(&id); err != nil {
		return 0, fmt.Errorf("QuestionRepository@Create | Error: %v", err)
	}

	return id, nil
}

func (r *QuestionRepository) Update(question models.Question) error {
	sql := `UPDATE questions SET sub_category_id = ?, question_text = ?, correct_answer = ?, acceptable_answer = ?, difficulty_level = ?, image_url = ?, updated_at = ? WHERE id = ?`

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("QuestionRepository@Update | Error: %v", err)
	}

	_, err = stmt.Exec(question.SubCategoryID, question.QuestionText, question.CorrectAnswer, question.AcceptableAnswer, question.DifficultyLevel, question.ImageURL, time.Now().Format(time.RFC3339), question.ID)

	if err != nil {
		return fmt.Errorf("QuestionRepository@Update | Error: %v", err)
	}

	return nil
}

func (r *QuestionRepository) Delete(id int) error {
	sql := `DELETE FROM questions WHERE id = ?`

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("QuestionRepository@Delete | Error: %v", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("QuestionRepository@Delete | Error: %v", err)
	}

	return nil
}
