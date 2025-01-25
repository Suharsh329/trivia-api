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

func TestGetQuestions(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	questionService := services.NewQuestionService(db)
	handler := NewQuestionHandler(questionService)

	expected := `{"data":[],"success":true}`

	req, err := http.NewRequest("GET", "/questions", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.GetQuestions(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateQuestion(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	questionService := services.NewQuestionService(db)
	handler := NewQuestionHandler(questionService)

	expected := `{"data":"Question created successfully","success":true}`

	question := models.Question{QuestionText: "Sample Question", CorrectAnswer: "Sample Answer", DifficultyLevel: 1}
	body, _ := json.Marshal(question)
	req, err := http.NewRequest("POST", "/questions", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.CreateQuestion(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUpdateQuestion(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	questionService := services.NewQuestionService(db)
	handler := NewQuestionHandler(questionService)

	expected := `{"data":"Question updated successfully","success":true}`

	question := models.Question{QuestionText: "Sample Question", CorrectAnswer: "Sample Answer", DifficultyLevel: 1}
	body, _ := json.Marshal(question)
	req, err := http.NewRequest("PUT", "/questions", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.UpdateQuestion(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDeleteQuestion(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	questionService := services.NewQuestionService(db)
	handler := NewQuestionHandler(questionService)

	expected := `{"data":"Question deleted successfully","success":true}`

	body, _ := json.Marshal(1)
	req, err := http.NewRequest("DELETE", "/questions", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.DeleteQuestion(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
