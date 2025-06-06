package services

import (
	"database/sql"
	"fmt"

	"trivia/internal/models"
	"trivia/internal/repositories"
)

type GameService struct {
	repo         repositories.GameRepository
	queueRepo    repositories.QuestionQueueRepository
	questionRepo repositories.QuestionRepository
}

func NewGameService(db *sql.DB) *GameService {
	gameRepo := repositories.NewGameRepository(db)
	queueRepo := repositories.NewQuestionQueueRepository(db)
	questionRepo := repositories.NewQuestionRepository(db)
	return &GameService{repo: *gameRepo, queueRepo: *queueRepo, questionRepo: *questionRepo}
}

func (s *GameService) isEmptyGame(game models.Game) bool {
	return game == (models.Game{}) || game.Name == ""
}

func (s *GameService) GetGames(filters map[string]any) ([]models.Game, error) {
	return s.repo.Get(filters)
}

func (s *GameService) CreateGame(game models.Game) (models.Game, error) {
	if s.isEmptyGame(game) {
		return models.Game{}, fmt.Errorf("GameService@CreateGame game cannot be empty")
	}
	return s.repo.Create(game)
}

func (s *GameService) UpdateGame(game models.Game) error {
	if s.isEmptyGame(game) {
		return fmt.Errorf("GameService@UpdateGame game cannot be empty")
	}
	return s.repo.Update(game)
}

func (s *GameService) DeleteGame(id int) error {
	if id <= 0 {
		return fmt.Errorf("GameService@DeleteGame invalid id")
	}
	return s.repo.Delete(id)
}

func (s *GameService) SetSelectedGame(gameId int, questionIds []int) error {
	for _, questionId := range questionIds {
		_, err := s.queueRepo.Add(gameId, questionId)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
	}

	return nil
}

func (s *GameService) SetRandomGame(gameId, numberOfQuestions int, percentages map[int]float64) error {
	questions := []models.Question{}
	numberOfQuestions += 10 // Add a buffer of 10 questions to ensure we have enough to work with
	for difficulty, percentage := range percentages {
		limit := int(float64(numberOfQuestions) * (percentage / 100.0))
		if limit <= 0 {
			continue
		}
		if limit > numberOfQuestions {
			limit = numberOfQuestions // Ensure we don't exceed the total number of questions requested
		}

		difficultyQuestions, err := s.questionRepo.Get(map[string]any{"difficulty_level": difficulty, "per_page": limit})
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		questions = append(questions, difficultyQuestions...)
	}

	for _, question := range questions {
		_, err := s.queueRepo.Add(gameId, question.ID)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
	}

	return nil
}

func (s *GameService) FetchQueueByGameId(gameId, limit int64) ([]models.QuestionQueueResponse, error) {
	queue, err := s.queueRepo.FetchQueueByGameId(gameId, limit)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return queue, nil
}
