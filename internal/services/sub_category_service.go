package services

import (
	"database/sql"
	"trivia/internal/models"
	"trivia/internal/repositories"
)

type SubCategoryService struct {
	repo repositories.SubCategoryRepository
}

func NewSubCategoryService(db *sql.DB) *SubCategoryService {
	subCategoryRepo := repositories.NewSubCategoryRepository(db)
	return &SubCategoryService{repo: *subCategoryRepo}
}

func (s *SubCategoryService) GetAllSubCategories(filters map[string]any) ([]models.SubCategory, error) {
	return s.repo.GetAll(filters)
}

func (s *SubCategoryService) GetSubCategoryById(id int) (models.SubCategory, error) {
	return s.repo.GetById(id)
}

func (s *SubCategoryService) CreateSubCategory(subCategory models.SubCategory) (int, error) {
	return s.repo.Create(subCategory)
}

func (s *SubCategoryService) UpdateSubCategory(subCategory models.SubCategory) error {
	return s.repo.Update(subCategory)
}

func (s *SubCategoryService) DeleteSubCategory(id int) error {
	return s.repo.Delete(id)
}
