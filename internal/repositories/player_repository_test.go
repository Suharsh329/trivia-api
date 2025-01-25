package repositories

import (
	"testing"

	"trivia/internal/models"
	"trivia/internal/utils"

	_ "github.com/mattn/go-sqlite3"
)

func TestPlayerGet(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewPlayerRepository(db)
	player := models.Player{
		Name:   "John Doe",
		GameID: 1,
	}

	numberOfPlayers := 1

	id, err := repo.Create(player)
	if err != nil {
		t.Fatalf("Failed to create player: %v", err)
	}

	// Act
	players, err := repo.Get(map[string]interface{}{"id": id})
	if err != nil {
		t.Fatalf("Failed to get player: %v", err)
	}

	// Assert
	if len(players) != 1 {
		t.Fatalf("Expected %v player, got %v", numberOfPlayers, len(players))
	}

	if players[0].ID != id {
		t.Fatalf("Expected player ID %d, got %d", id, players[0].ID)
	}
}

func TestPlayerCreate(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewPlayerRepository(db)

	player := models.Player{
		Name:   "John Doe",
		GameID: 1,
	}

	// Act
	id, err := repo.Create(player)
	if err != nil {
		t.Fatalf("Failed to create player: %v", err)
	}

	// Assert
	if id == 0 {
		t.Fatalf("Expected valid player ID, got %d", id)
	}
}

func TestPlayerUpdate(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewPlayerRepository(db)

	player := models.Player{
		Name:   "John Doe",
		GameID: 2,
	}

	id, err := repo.Create(player)
	if err != nil {
		t.Fatalf("Failed to create player: %v", err)
	}

	player.ID = id
	player.Name = "Jane Doe"

	// Act
	err = repo.Update(player)
	if err != nil {
		t.Fatalf("Failed to update player: %v", err)
	}

	// Assert
	players, err := repo.Get(map[string]any{"id": id})
	if err != nil {
		t.Fatalf("Failed to get player: %v", err)
	}

	if players[0].Name != player.Name {
		t.Fatalf("Expected player name %s, got %s", player.Name, players[0].Name)
	}
}

func TestPlayerDelete(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewPlayerRepository(db)

	player := models.Player{
		Name:   "John Doe",
		GameID: 1,
	}

	id, err := repo.Create(player)
	if err != nil {
		t.Fatalf("Failed to create player: %v", err)
	}

	// Act
	err = repo.Delete(id)
	if err != nil {
		t.Fatalf("Failed to delete player: %v", err)
	}

	// Assert
	players, err := repo.Get(map[string]interface{}{"id": id})
	if err != nil {
		t.Fatalf("Failed to get player: %v", err)
	}

	if len(players) != 0 {
		t.Fatalf("Expected 0 players, got %v", len(players))
	}
}
