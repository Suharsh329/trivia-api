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

func TestGetCategories(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	categoryService := services.NewCategoryService(db)
	handler := NewCategoryHandler(categoryService)

	expected := `{"data":null,"success":true}`

	req, err := http.NewRequest("GET", "/categories", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.GetCategories(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateCategory(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	categoryService := services.NewCategoryService(db)
	handler := NewCategoryHandler(categoryService)

	expected := `{"data":"Category created successfully","success":true}`

	category := models.Category{Name: "Sample Category"}
	body, _ := json.Marshal(category)
	req, err := http.NewRequest("POST", "/categories", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.CreateCategory(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUpdateCategory(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	categoryService := services.NewCategoryService(db)
	handler := NewCategoryHandler(categoryService)

	expected := `{"data":"Category updated successfully","success":true}`

	category := models.Category{Name: "Sample Category"}
	body, _ := json.Marshal(category)
	req, err := http.NewRequest("PUT", "/categories", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.UpdateCategory(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDeleteCategory(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	categoryService := services.NewCategoryService(db)
	handler := NewCategoryHandler(categoryService)

	expected := `{"data":"Category deleted successfully","success":true}`

	body, _ := json.Marshal(1)
	req, err := http.NewRequest("DELETE", "/categories", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.DeleteCategory(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
