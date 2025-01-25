package services

import (
	"database/sql"
	"trivia/internal/models"
	"trivia/internal/repositories"
)

type CategoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(db *sql.DB) *CategoryService {
	categoryRepo := repositories.NewCategoryRepository(db)
	return &CategoryService{repo: *categoryRepo}
}

func (s *CategoryService) GetAllCategories(filters map[string]any) ([]models.Category, error) {
	return s.repo.GetAll(filters)
}

func (s *CategoryService) GetCategoryById(id int) (models.Category, error) {
	return s.repo.GetById(id)
}

func (s *CategoryService) CreateCategory(category models.Category) (int, error) {
	return s.repo.Create(category)
}

func (s *CategoryService) UpdateCategory(category models.Category) error {
	return s.repo.Update(category)
}

func (s *CategoryService) DeleteCategory(id int) error {
	return s.repo.Delete(id)
}
