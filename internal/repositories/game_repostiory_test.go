package repositories

import (
	"testing"

	"trivia/internal/models"
	"trivia/internal/utils"

	_ "github.com/mattn/go-sqlite3"
)

func TestGameGet(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewGameRepository(db)
	game := models.Game{
		Name: "Trivia Game",
	}

	numberOfGames := 1

	newGame, err := repo.Create(game)
	if err != nil {
		t.Fatalf("Failed to create game: %v", err)
	}

	// Act
	games, err := repo.Get(map[string]interface{}{"id": newGame.ID})
	if err != nil {
		t.Fatalf("Failed to get game: %v", err)
	}

	// Assert
	if len(games) != 1 {
		t.Fatalf("Expected %v game, got %v", numberOfGames, len(games))
	}

	if games[0].ID != newGame.ID {
		t.Fatalf("Expected game ID %d, got %d", newGame.ID, games[0].ID)
	}
}

func TestGameCreate(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewGameRepository(db)

	game := models.Game{
		Name: "Trivia Game",
	}

	// Act
	newGame, err := repo.Create(game)
	if err != nil {
		t.Fatalf("Failed to create game: %v", err)
	}

	// Assert
	if newGame.ID == 0 {
		t.Fatalf("Expected valid game ID, got %d", newGame.ID)
	}
}

func TestGameUpdate(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewGameRepository(db)

	game := models.Game{
		Name: "Trivia Game",
	}

	newGame, err := repo.Create(game)
	if err != nil {
		t.Fatalf("Failed to create game: %v", err)
	}

	game.ID = newGame.ID
	game.Name = "Updated Trivia Game"

	// Act
	err = repo.Update(game)
	if err != nil {
		t.Fatalf("Failed to update game: %v", err)
	}

	// Assert
	games, err := repo.Get(map[string]any{"id": newGame.ID})
	if err != nil {
		t.Fatalf("Failed to get game: %v", err)
	}

	if games[0].Name != game.Name {
		t.Fatalf("Expected game name %s, got %s", game.Name, games[0].Name)
	}
}

func TestGameDelete(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewGameRepository(db)

	game := models.Game{
		Name: "Trivia Game",
	}

	newGame, err := repo.Create(game)
	if err != nil {
		t.Fatalf("Failed to create game: %v", err)
	}

	// Act
	err = repo.Delete(newGame.ID)
	if err != nil {
		t.Fatalf("Failed to delete game: %v", err)
	}

	// Assert
	games, err := repo.Get(map[string]interface{}{"id": newGame.ID})
	if err != nil {
		t.Fatalf("Failed to get game: %v", err)
	}

	if len(games) != 0 {
		t.Fatalf("Expected 0 games, got %v", len(games))
	}
}
