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

	// Parse the response to verify structure
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	// Check success field
	if success, ok := response["success"].(bool); !ok || !success {
		t.Errorf("Expected success to be true, got %v", response["success"])
	}

	// Check data structure
	data, ok := response["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("Expected data to be an object, got %T", response["data"])
	}

	// Check message
	if message, ok := data["message"].(string); !ok || message != "Game created successfully" {
		t.Errorf("Expected message 'Game created successfully', got %v", data["message"])
	}

	// Check game object exists
	gameData, ok := data["game"].(map[string]interface{})
	if !ok {
		t.Fatalf("Expected game object in data, got %T", data["game"])
	}

	// Check game name
	if gameName, ok := gameData["game_name"].(string); !ok || gameName != "Sample Game" {
		t.Errorf("Expected game name 'Sample Game', got %v", gameData["game_name"])
	}

	// Check game ID exists and is positive
	if gameID, ok := gameData["id"].(float64); !ok || gameID <= 0 {
		t.Errorf("Expected positive game ID, got %v", gameData["id"])
	}
}

func TestUpdateGame(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	gameService := services.NewGameService(db)
	handler := NewGameHandler(gameService)

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

	// Parse the response to verify structure
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	// Check success field
	if success, ok := response["success"].(bool); !ok || !success {
		t.Errorf("Expected success to be true, got %v", response["success"])
	}

	// Check data message
	if data, ok := response["data"].(string); !ok || data != "Game updated successfully" {
		t.Errorf("Expected data 'Game updated successfully', got %v", response["data"])
	}
}

func TestDeleteGame(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	gameService := services.NewGameService(db)
	handler := NewGameHandler(gameService)

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

	// Parse the response to verify structure
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	// Check success field
	if success, ok := response["success"].(bool); !ok || !success {
		t.Errorf("Expected success to be true, got %v", response["success"])
	}

	// Check data message
	if data, ok := response["data"].(string); !ok || data != "Game deleted successfully" {
		t.Errorf("Expected data 'Game deleted successfully', got %v", response["data"])
	}
}
func TestSetSelectedGame(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	gameService := services.NewGameService(db)
	handler := NewGameHandler(gameService)

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

	// Parse the response to verify structure
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	// Check success field
	if success, ok := response["success"].(bool); !ok || !success {
		t.Errorf("Expected success to be true, got %v", response["success"])
	}

	// Check data message
	if data, ok := response["data"].(string); !ok || data != "Selected game set successfully" {
		t.Errorf("Expected data 'Selected game set successfully', got %v", response["data"])
	}
}

func TestSetRandomGame(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	gameService := services.NewGameService(db)
	handler := NewGameHandler(gameService)

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

	// Parse the response to verify structure
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	// Check success field
	if success, ok := response["success"].(bool); !ok || !success {
		t.Errorf("Expected success to be true, got %v", response["success"])
	}

	// Check data message
	if data, ok := response["data"].(string); !ok || data != "Random game set successfully" {
		t.Errorf("Expected data 'Random game set successfully', got %v", response["data"])
	}
}
