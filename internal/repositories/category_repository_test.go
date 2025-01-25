package repositories

import (
	"testing"
	"trivia/internal/models"
	"trivia/internal/utils"

	_ "github.com/mattn/go-sqlite3"
)

func TestCategoryGetAll(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewCategoryRepository(db)
	category := models.Category{
		Name:     "Category A",
		ImageUrl: "http://example.com/image.png",
	}

	id, err := repo.Create(category)
	if err != nil {
		t.Fatalf("Failed to create category: %v", err)
	}

	// Act
	categories, err := repo.GetAll(map[string]interface{}{"id": id})
	if err != nil {
		t.Fatalf("Failed to get category: %v", err)
	}

	// Assert
	if len(categories) != 1 {
		t.Fatalf("Expected 1 category, got %v", len(categories))
	}

	if categories[0].ID != id {
		t.Fatalf("Expected category ID %d, got %d", id, categories[0].ID)
	}
}

func TestCategoryCreate(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewCategoryRepository(db)

	category := models.Category{
		Name:     "Category A",
		ImageUrl: "http://example.com/image.png",
	}

	// Act
	id, err := repo.Create(category)
	if err != nil {
		t.Fatalf("Failed to create category: %v", err)
	}

	// Assert
	if id == 0 {
		t.Fatalf("Expected valid category ID, got %d", id)
	}
}

func TestCategoryUpdate(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewCategoryRepository(db)

	category := models.Category{
		Name:     "Category A",
		ImageUrl: "http://example.com/image.png",
	}

	id, err := repo.Create(category)
	if err != nil {
		t.Fatalf("Failed to create category: %v", err)
	}

	category.ID = id
	category.Name = "Category B"
	category.ImageUrl = "http://example.com/new_image.png"

	// Act
	err = repo.Update(category)
	if err != nil {
		t.Fatalf("Failed to update category: %v", err)
	}

	// Assert
	categories, err := repo.GetAll(map[string]any{"id": id})
	if err != nil {
		t.Fatalf("Failed to get category: %v", err)
	}

	if categories[0].Name != category.Name {
		t.Fatalf("Expected category name %s, got %s", category.Name, categories[0].Name)
	}
}

func TestCategoryDelete(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewCategoryRepository(db)

	category := models.Category{
		Name:     "Category A",
		ImageUrl: "http://example.com/image.png",
	}

	id, err := repo.Create(category)
	if err != nil {
		t.Fatalf("Failed to create category: %v", err)
	}

	// Act
	err = repo.Delete(id)
	if err != nil {
		t.Fatalf("Failed to delete category: %v", err)
	}

	// Assert
	categories, err := repo.GetAll(map[string]interface{}{"id": id})
	if err != nil {
		t.Fatalf("Failed to get category: %v", err)
	}

	if len(categories) != 0 {
		t.Fatalf("Expected 0 categories, got %v", len(categories))
	}
}
