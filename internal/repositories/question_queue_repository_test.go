package repositories

import (
	"database/sql"
	"testing"
	"trivia/internal/models"
	"trivia/internal/utils"

	_ "github.com/mattn/go-sqlite3"
)

func TestQuestionQueueFetchQueueByGameId(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewQuestionQueueRepository(db)

	gameId := 1
	question := models.Question{
		ID:               1,
		QuestionText:     "What is the capital of France?",
		CorrectAnswer:    "Paris",
		AcceptableAnswer: sql.NullString{String: "", Valid: true},
		SubCategoryID:    1,
		DifficultyLevel:  1,
	}
	_, err := db.Exec("INSERT INTO questions (id, question_text, correct_answer, acceptable_answer, sub_category_id, difficulty_level) VALUES (?, ?, ?, ?, ?, ?)", question.ID, question.QuestionText, question.CorrectAnswer, question.AcceptableAnswer, question.SubCategoryID, question.DifficultyLevel)
	if err != nil {
		t.Fatalf("Failed to add question: %v", err)
	}
	subcategory := models.SubCategory{
		ID:   1,
		Name: "Geography",
	}
	_, err = db.Exec("INSERT INTO sub_categories (id, name) VALUES (?, ?)", subcategory.ID, subcategory.Name)
	if err != nil {
		t.Fatalf("Failed to add subcategory: %v", err)
	}

	_, err = repo.Add(gameId, 1)
	if err != nil {
		t.Fatalf("Failed to add question to queue: %v", err)
	}

	// Act
	queue, err := repo.FetchQueueByGameId(gameId, 1)
	if err != nil {
		t.Fatalf("Failed to fetch question queue: %v", err)
	}

	// Assert
	if queue[0].GameID != gameId {
		t.Fatalf("Expected question ID %d, got %d", gameId, queue[0].GameID)
	}
}

func TestQuestionQueueAdd(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewQuestionQueueRepository(db)

	// Act
	id, err := repo.Add(1, 1)
	if err != nil {
		t.Fatalf("Failed to add question to queue: %v", err)
	}

	// Assert
	if id == 0 {
		t.Fatalf("Expected valid queue ID, got %d", id)
	}
}

func TestQuestionQueueDelete(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewQuestionQueueRepository(db)
	question := models.Question{
		ID:               1,
		QuestionText:     "What is the capital of France?",
		CorrectAnswer:    "Paris",
		AcceptableAnswer: sql.NullString{String: "", Valid: true},
		SubCategoryID:    1,
		DifficultyLevel:  1,
	}
	_, err := db.Exec("INSERT INTO questions (id, question_text, correct_answer, acceptable_answer, sub_category_id, difficulty_level) VALUES (?, ?, ?, ?, ?, ?)", question.ID, question.QuestionText, question.CorrectAnswer, question.AcceptableAnswer, question.SubCategoryID, question.DifficultyLevel)
	if err != nil {
		t.Fatalf("Failed to add question: %v", err)
	}
	subcategory := models.SubCategory{
		ID:   1,
		Name: "Geography",
	}
	_, err = db.Exec("INSERT INTO sub_categories (id, name) VALUES (?, ?)", subcategory.ID, subcategory.Name)
	if err != nil {
		t.Fatalf("Failed to add subcategory: %v", err)
	}

	id, err := repo.Add(1, 1)
	if err != nil {
		t.Fatalf("Failed to add question to queue: %v", err)
	}

	// Act
	err = repo.Delete(id)
	if err != nil {
		t.Fatalf("Failed to delete question from queue: %v", err)
	}

	// Assert
	queue, err := repo.FetchQueueByGameId(1, 1)
	if err != nil {
		t.Fatalf("Failed to fetch question queue: %v", err)
	}

	if len(queue) != 0 {
		t.Fatalf("Expected question ID to be deleted, but found %d", len(queue))
	}
}
