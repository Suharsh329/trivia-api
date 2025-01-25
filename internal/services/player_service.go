package services

import (
	"database/sql"
	"trivia/internal/models"
	"trivia/internal/repositories"
)

type PlayerService struct {
	repo repositories.PlayerRepository
}

func (s *PlayerService) GetAllPlayers(filters map[string]any) ([]models.Player, error) {
	return s.repo.Get(filters)
}

func NewPlayerService(db *sql.DB) *PlayerService {
	playerRepo := repositories.NewPlayerRepository(db)
	return &PlayerService{repo: *playerRepo}
}

func (s *PlayerService) CreatePlayer(player models.Player) (int, error) {
	return s.repo.Create(player)
}

func (s *PlayerService) UpdatePlayer(player models.Player) error {
	return s.repo.Update(player)
}

func (s *PlayerService) DeletePlayer(id int) error {
	return s.repo.Delete(id)
}
