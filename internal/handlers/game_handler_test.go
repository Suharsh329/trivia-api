package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"trivia/internal/models"
	"trivia/internal/services"
	"trivia/internal/utils"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetGames(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	gameService := services.NewGameService(db)
	handler := NewGameHandler(gameService)

	expected := `{"data":[],"success":true}`

	req, err := http.NewRequest("GET", "/games", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.GetGames(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateGame(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	gameService := services.NewGameService(db)
	handler := NewGameHandler(gameService)

	expected := `{"data":"Game created successfully","success":true}`

	game := models.Game{Name: "Sample Game"}
	body, _ := json.Marshal(game)
	req, err := http.NewRequest("POST", "/games", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.CreateGame(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUpdateGame(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	gameService := services.NewGameService(db)
	handler := NewGameHandler(gameService)

	expected := `{"data":"Game updated successfully","success":true}`

	game := models.Game{Name: "Sample Game"}
	body, _ := json.Marshal(game)
	req, err := http.NewRequest("PUT", "/games", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.UpdateGame(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDeleteGame(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	gameService := services.NewGameService(db)
	handler := NewGameHandler(gameService)

	expected := `{"data":"Game deleted successfully","success":true}`

	body, _ := json.Marshal(1)
	req, err := http.NewRequest("DELETE", "/games", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.DeleteGame(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
func TestSetSelectedGame(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	gameService := services.NewGameService(db)
	handler := NewGameHandler(gameService)

	expected := `{"data":"Selected game set successfully","success":true}`

	requestData := struct {
		GameId      int   `json:"gameId"`
		QuestionIds []int `json:"questionIds"`
	}{
		GameId:      1,
		QuestionIds: []int{1, 2, 3},
	}
	body, _ := json.Marshal(requestData)
	req, err := http.NewRequest("POST", "/games/set-selected-game", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.SetSelectedGame(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestSetRandomGame(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	gameService := services.NewGameService(db)
	handler := NewGameHandler(gameService)

	expected := `{"data":"Random game set successfully","success":true}`

	requestData := struct {
		GameId            int             `json:"gameId"`
		NumberOfQuestions int             `json:"numberOfQuestions"`
		Percentages       map[int]float64 `json:"percentages"`
	}{
		GameId:            1,
		NumberOfQuestions: 10,
		Percentages:       map[int]float64{1: 0.5, 2: 0.3, 3: 0.2},
	}
	body, _ := json.Marshal(requestData)
	req, err := http.NewRequest("POST", "/games/set-random-game", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.SetRandomGame(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
