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

func TestGetPlayers(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	playerService := services.NewPlayerService(db)
	handler := NewPlayerHandler(playerService)

	expected := `{"data":[],"success":true}`

	req, err := http.NewRequest("GET", "/players", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.GetPlayers(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreatePlayer(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	playerService := services.NewPlayerService(db)
	handler := NewPlayerHandler(playerService)

	expected := `{"data":"Player created successfully","success":true}`

	player := models.Player{Name: "Sample Player"}
	body, _ := json.Marshal(player)
	req, err := http.NewRequest("POST", "/players", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.CreatePlayer(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUpdatePlayer(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	playerService := services.NewPlayerService(db)
	handler := NewPlayerHandler(playerService)

	expected := `{"data":"Player updated successfully","success":true}`

	player := models.Player{Name: "Sample Player"}
	body, _ := json.Marshal(player)
	req, err := http.NewRequest("PUT", "/players", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.UpdatePlayer(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDeletePlayer(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	playerService := services.NewPlayerService(db)
	handler := NewPlayerHandler(playerService)

	expected := `{"data":"Player deleted successfully","success":true}`

	body, _ := json.Marshal(1)
	req, err := http.NewRequest("DELETE", "/players", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.DeletePlayer(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
