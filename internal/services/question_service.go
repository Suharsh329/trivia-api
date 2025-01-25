package services

import (
	"database/sql"
	"fmt"

	"trivia/internal/models"
	"trivia/internal/repositories"
)

type QuestionService struct {
	repo repositories.QuestionRepository
}

func NewQuestionService(db *sql.DB) *QuestionService {
	questionRepo := repositories.NewQuestionRepository(db)
	return &QuestionService{repo: *questionRepo}
}

func (s *QuestionService) isEmptyQuestion(question models.Question) bool {
	return question == (models.Question{}) ||
		question.QuestionText == "" ||
		question.CorrectAnswer == "" ||
		question.DifficultyLevel <= 0
}

func (s *QuestionService) GetQuestions(filters map[string]any) ([]models.Question, error) {
	return s.repo.Get(filters)
}

func (s *QuestionService) CreateQuestion(question models.Question) (int, error) {
	if s.isEmptyQuestion(question) {
		return 0, fmt.Errorf("QuestionService@CreateQuestion question cannot be empty")
	}
	return s.repo.Create(question)
}

func (s *QuestionService) UpdateQuestion(question models.Question) error {
	if s.isEmptyQuestion(question) {
		return fmt.Errorf("QuestionService@UpdateQuestion question cannot be empty")
	}
	return s.repo.Update(question)
}

func (s *QuestionService) DeleteQuestion(id int) error {
	if id <= 0 {
		return fmt.Errorf("QuestionService@DeleteQuestion invalid id")
	}
	return s.repo.Delete(id)
}
