package handlers

import (
	"encoding/json"
	"net/http"
	"trivia/internal/models"
	"trivia/internal/response"
	"trivia/internal/services"
)

type CategoryHandler struct {
	Service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: service}
}

func ValidateCategory(category models.Category) bool {
	return category.Name != ""
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	filters := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&filters)

	categories, err := h.Service.GetAllCategories(filters)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}

	response.Success(w, categories)
}

func (h *CategoryHandler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	var id int
	json.NewDecoder(r.Body).Decode(&id)

	categories, err := h.Service.GetCategoryById(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch category")
		return
	}

	response.Success(w, categories)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	json.NewDecoder(r.Body).Decode(&category)

	if !ValidateCategory(category) {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	_, err := h.Service.CreateCategory(category)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create a category")
		return
	}

	response.Success(w, "Category created successfully")
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.UpdateCategory(category)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update the category")
		return
	}

	response.Success(w, "Category updated successfully")
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	var id int

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.DeleteCategory(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete the category")
		return
	}

	response.Success(w, "Category deleted successfully")
}
