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

func TestGetTeams(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	teamService := services.NewTeamService(db)
	handler := NewTeamHandler(teamService)

	expected := `{"data":[],"success":true}`

	req, err := http.NewRequest("GET", "/teams", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.GetTeams(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateTeam(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	teamService := services.NewTeamService(db)
	handler := NewTeamHandler(teamService)

	expected := `{"data":"Team created successfully","success":true}`

	team := models.Team{Name: "Sample Team"}
	body, _ := json.Marshal(team)
	req, err := http.NewRequest("POST", "/teams", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.CreateTeam(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUpdateTeam(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	teamService := services.NewTeamService(db)
	handler := NewTeamHandler(teamService)

	expected := `{"data":"Team updated successfully","success":true}`

	team := models.Team{Name: "Sample Team"}
	body, _ := json.Marshal(team)
	req, err := http.NewRequest("PUT", "/teams", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.UpdateTeam(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDeleteTeam(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	teamService := services.NewTeamService(db)
	handler := NewTeamHandler(teamService)

	expected := `{"data":"Team deleted successfully","success":true}`

	body, _ := json.Marshal(1)
	req, err := http.NewRequest("DELETE", "/teams", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.DeleteTeam(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
