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

func TestGetSubCategories(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	subCategoryService := services.NewSubCategoryService(db)
	handler := NewSubCategoryHandler(subCategoryService)

	expected := `{"data":null,"success":true}`

	req, err := http.NewRequest("GET", "/sub_categories", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.GetSubCategories(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateSubCategory(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	subCategoryService := services.NewSubCategoryService(db)
	handler := NewSubCategoryHandler(subCategoryService)

	expected := `{"data":"Sub-category created successfully","success":true}`

	subCategory := models.SubCategory{Name: "Sample SubCategory"}
	body, _ := json.Marshal(subCategory)
	req, err := http.NewRequest("POST", "/sub_categories", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.CreateSubCategory(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUpdateSubCategory(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	subCategoryService := services.NewSubCategoryService(db)
	handler := NewSubCategoryHandler(subCategoryService)

	expected := `{"data":"Sub-category updated successfully","success":true}`

	subCategory := models.SubCategory{Name: "Sample SubCategory"}
	body, _ := json.Marshal(subCategory)
	req, err := http.NewRequest("PUT", "/sub_categories", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.UpdateSubCategory(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDeleteSubCategory(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	subCategoryService := services.NewSubCategoryService(db)
	handler := NewSubCategoryHandler(subCategoryService)

	expected := `{"data":"Sub-category deleted successfully","success":true}`

	body, _ := json.Marshal(1)
	req, err := http.NewRequest("DELETE", "/sub_categories", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler.DeleteSubCategory(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
