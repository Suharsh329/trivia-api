package repositories

import (
	"database/sql"
	"testing"

	"trivia/internal/models"
	"trivia/internal/utils"

	_ "github.com/mattn/go-sqlite3"
)

func TestQuestionGet(t *testing.T) {
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewQuestionRepository(db)

	question := models.Question{
		SubCategoryID:   1,
		QuestionText:    "What is the capital of France?",
		CorrectAnswer:   "Paris",
		DifficultyLevel: 1,
	}

	id, err := repo.Create(question)
	if err != nil {
		t.Fatalf("Failed to create question: %v", err)
	}

	questions, err := repo.Get(map[string]any{"id": id})
	if err != nil {
		t.Fatalf("Failed to get question: %v", err)
	}

	if len(questions) != 1 {
		t.Fatalf("Expected 1 question, got 0")
	}

	if questions[0].ID != id {
		t.Fatalf("Expected question ID %d, got %d", id, questions[0].ID)
	}
}

func TestQuestionCreate(t *testing.T) {
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewQuestionRepository(db)

	question := models.Question{
		SubCategoryID:    1,
		QuestionText:     "What is the capital of France?",
		CorrectAnswer:    "Paris",
		AcceptableAnswer: sql.NullString{String: "paris", Valid: true},
		DifficultyLevel:  1,
	}

	id, err := repo.Create(question)
	if err != nil {
		t.Fatalf("Failed to create question: %v", err)
	}

	if id == 0 {
		t.Fatalf("Expected valid question ID, got %d", id)
	}
}

func TestQuestionUpdate(t *testing.T) {
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewQuestionRepository(db)

	question := models.Question{
		SubCategoryID:   1,
		QuestionText:    "What is the capital of France?",
		CorrectAnswer:   "Paris",
		DifficultyLevel: 1,
	}

	id, err := repo.Create(question)
	if err != nil {
		t.Fatalf("Failed to create question: %v", err)
	}

	question.ID = id
	question.SubCategoryID = 2
	question.QuestionText = "What is the capital of Germany?"
	question.CorrectAnswer = "Berlin"
	question.AcceptableAnswer = sql.NullString{String: "berlin", Valid: true}
	question.DifficultyLevel = 2

	err = repo.Update(question)
	if err != nil {
		t.Fatalf("Failed to update question: %v", err)
	}
}

func TestQuestionDelete(t *testing.T) {
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewQuestionRepository(db)

	question := models.Question{
		SubCategoryID:   1,
		QuestionText:    "What is the capital of France?",
		CorrectAnswer:   "Paris",
		DifficultyLevel: 1,
	}

	id, err := repo.Create(question)
	if err != nil {
		t.Fatalf("Failed to create question: %v", err)
	}

	err = repo.Delete(id)
	if err != nil {
		t.Fatalf("Failed to delete question: %v", err)
	}
}
