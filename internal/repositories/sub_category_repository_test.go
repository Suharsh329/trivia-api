package repositories

import (
	"database/sql"
	"testing"
	"trivia/internal/models"
	"trivia/internal/utils"

	_ "github.com/mattn/go-sqlite3"
)

func TestSubCategoryGetAll(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewSubCategoryRepository(db)
	subCategory := models.SubCategory{
		Name:       "SubCategory A",
		CategoryID: 1,
		ImageUrl:   sql.NullString{String: "http://example.com/image.png", Valid: true},
	}

	id, err := repo.Create(subCategory)
	if err != nil {
		t.Fatalf("Failed to create sub-category: %v", err)
	}

	// Act
	subCategories, err := repo.GetAll(map[string]interface{}{"id": id})
	if err != nil {
		t.Fatalf("Failed to get sub-category: %v", err)
	}

	// Assert
	if len(subCategories) != 1 {
		t.Fatalf("Expected 1 sub-category, got %v", len(subCategories))
	}

	if subCategories[0].ID != id {
		t.Fatalf("Expected sub-category ID %d, got %d", id, subCategories[0].ID)
	}
}

func TestSubCategoryCreate(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewSubCategoryRepository(db)

	subCategory := models.SubCategory{
		Name:       "SubCategory A",
		CategoryID: 1,
		ImageUrl:   sql.NullString{String: "http://example.com/image.png", Valid: true},
	}

	// Act
	id, err := repo.Create(subCategory)
	if err != nil {
		t.Fatalf("Failed to create sub-category: %v", err)
	}

	// Assert
	if id == 0 {
		t.Fatalf("Expected valid sub-category ID, got %d", id)
	}
}

func TestSubCategoryUpdate(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewSubCategoryRepository(db)

	subCategory := models.SubCategory{
		Name:       "SubCategory A",
		CategoryID: 1,
		ImageUrl:   sql.NullString{String: "http://example.com/image.png", Valid: true},
	}

	id, err := repo.Create(subCategory)
	if err != nil {
		t.Fatalf("Failed to create sub-category: %v", err)
	}

	subCategory.ID = id
	subCategory.Name = "SubCategory B"
	subCategory.ImageUrl = sql.NullString{String: "http://example.com/image.png", Valid: true}

	// Act
	err = repo.Update(subCategory)
	if err != nil {
		t.Fatalf("Failed to update sub-category: %v", err)
	}

	// Assert
	subCategories, err := repo.GetAll(map[string]any{"id": id})
	if err != nil {
		t.Fatalf("Failed to get sub-category: %v", err)
	}

	if subCategories[0].Name != subCategory.Name {
		t.Fatalf("Expected sub-category name %s, got %s", subCategory.Name, subCategories[0].Name)
	}
}

func TestSubCategoryDelete(t *testing.T) {
	// Arrange
	db := utils.SetupTestDB()
	defer db.Close()

	repo := NewSubCategoryRepository(db)

	subCategory := models.SubCategory{
		Name:       "SubCategory A",
		CategoryID: 1,
		ImageUrl:   sql.NullString{String: "http://example.com/image.png", Valid: true},
	}

	id, err := repo.Create(subCategory)
	if err != nil {
		t.Fatalf("Failed to create sub-category: %v", err)
	}

	// Act
	err = repo.Delete(id)
	if err != nil {
		t.Fatalf("Failed to delete sub-category: %v", err)
	}

	// Assert
	subCategories, err := repo.GetAll(map[string]interface{}{"id": id})
	if err != nil {
		t.Fatalf("Failed to get sub-category: %v", err)
	}

	if len(subCategories) != 0 {
		t.Fatalf("Expected 0 sub-categories, got %v", len(subCategories))
	}
}
