package services

import (
	"database/sql"
	"trivia/internal/models"
	"trivia/internal/repositories"
)

type TeamService struct {
	repo repositories.TeamRepository
}

func NewTeamService(db *sql.DB) *TeamService {
	teamRepo := repositories.NewTeamRepository(db)
	return &TeamService{repo: *teamRepo}
}

func (s *TeamService) GetAllTeams(filters map[string]any) ([]models.Team, error) {
	return s.repo.Get(filters)
}

func (s *TeamService) CreateTeam(team models.Team) (int, error) {
	return s.repo.Create(team)
}

func (s *TeamService) UpdateTeam(team models.Team) error {
	return s.repo.Update(team)
}

func (s *TeamService) DeleteTeam(id int) error {
	return s.repo.Delete(id)
}
