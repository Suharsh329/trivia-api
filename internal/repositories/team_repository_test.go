package repositories

import (
	"testing"

	"trivia/internal/models"
	"trivia/internal/utils"

	_ "github.com/mattn/go-sqlite3"
)

func TestTeamGet(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewTeamRepository(db)
	team := models.Team{
		Name:   "Team A",
		GameID: 1,
	}

	numberOfTeams := 1

	id, err := repo.Create(team)
	if err != nil {
		t.Fatalf("Failed to create team: %v", err)
	}

	// Act
	teams, err := repo.Get(map[string]interface{}{"id": id})
	if err != nil {
		t.Fatalf("Failed to get team: %v", err)
	}

	// Assert
	if len(teams) != 1 {
		t.Fatalf("Expected %v team, got %v", numberOfTeams, len(teams))
	}

	if teams[0].ID != id {
		t.Fatalf("Expected team ID %d, got %d", id, teams[0].ID)
	}
}

func TestTeamCreate(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewTeamRepository(db)

	team := models.Team{
		Name:   "Team A",
		GameID: 1,
	}

	// Act
	id, err := repo.Create(team)
	if err != nil {
		t.Fatalf("Failed to create team: %v", err)
	}

	// Assert
	if id == 0 {
		t.Fatalf("Expected valid team ID, got %d", id)
	}
}

func TestTeamUpdate(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewTeamRepository(db)

	team := models.Team{
		Name:   "Team A",
		GameID: 1,
	}

	id, err := repo.Create(team)
	if err != nil {
		t.Fatalf("Failed to create team: %v", err)
	}

	team.ID = id
	team.Name = "Team B"

	// Act
	err = repo.Update(team)
	if err != nil {
		t.Fatalf("Failed to update team: %v", err)
	}

	// Assert
	teams, err := repo.Get(map[string]any{"id": id})
	if err != nil {
		t.Fatalf("Failed to get team: %v", err)
	}

	if teams[0].Name != team.Name {
		t.Fatalf("Expected team name %s, got %s", team.Name, teams[0].Name)
	}
}

func TestTeamDelete(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewTeamRepository(db)

	team := models.Team{
		Name:   "Team A",
		GameID: 1,
	}

	id, err := repo.Create(team)
	if err != nil {
		t.Fatalf("Failed to create team: %v", err)
	}

	// Act
	err = repo.Delete(id)
	if err != nil {
		t.Fatalf("Failed to delete team: %v", err)
	}

	// Assert
	teams, err := repo.Get(map[string]interface{}{"id": id})
	if err != nil {
		t.Fatalf("Failed to get team: %v", err)
	}

	if len(teams) != 0 {
		t.Fatalf("Expected 0 teams, got %v", len(teams))
	}
}
