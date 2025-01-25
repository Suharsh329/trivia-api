package handlers

import (
	"encoding/json"
	"net/http"
	"trivia/internal/models"
	"trivia/internal/response"
	"trivia/internal/services"
)

type SubCategoryHandler struct {
	Service *services.SubCategoryService
}

func NewSubCategoryHandler(service *services.SubCategoryService) *SubCategoryHandler {
	return &SubCategoryHandler{Service: service}
}

func ValidateSubCategory(subCategory models.SubCategory) bool {
	return subCategory.Name != ""
}

func (h *SubCategoryHandler) GetSubCategories(w http.ResponseWriter, r *http.Request) {
	filters := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&filters)

	subCategories, err := h.Service.GetAllSubCategories(filters)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch sub-categories")
		return
	}

	response.Success(w, subCategories)
}

func (h *SubCategoryHandler) GetSubCategoryById(w http.ResponseWriter, r *http.Request) {
	var id int
	json.NewDecoder(r.Body).Decode(&id)

	subCategories, err := h.Service.GetSubCategoryById(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch sub-category")
		return
	}

	response.Success(w, subCategories)
}

func (h *SubCategoryHandler) CreateSubCategory(w http.ResponseWriter, r *http.Request) {
	var subCategory models.SubCategory

	json.NewDecoder(r.Body).Decode(&subCategory)

	if !ValidateSubCategory(subCategory) {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	_, err := h.Service.CreateSubCategory(subCategory)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create a sub-category")
		return
	}

	response.Success(w, "Sub-category created successfully")
}

func (h *SubCategoryHandler) UpdateSubCategory(w http.ResponseWriter, r *http.Request) {
	var subCategory models.SubCategory

	err := json.NewDecoder(r.Body).Decode(&subCategory)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.UpdateSubCategory(subCategory)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update the sub-category")
		return
	}

	response.Success(w, "Sub-category updated successfully")
}

func (h *SubCategoryHandler) DeleteSubCategory(w http.ResponseWriter, r *http.Request) {
	var id int

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.DeleteSubCategory(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete the sub-category")
		return
	}

	response.Success(w, "Sub-category deleted successfully")
}
