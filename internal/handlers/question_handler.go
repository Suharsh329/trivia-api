package handlers

import (
	"encoding/json"
	"net/http"
	"trivia/internal/models"
	"trivia/internal/response"
	"trivia/internal/services"
)

type QuestionHandler struct {
	Service *services.QuestionService
}

func NewQuestionHandler(service *services.QuestionService) *QuestionHandler {
	return &QuestionHandler{Service: service}
}

func ValidateQuestion(question models.Question) bool {
	return !(question.QuestionText == "" ||
		question.CorrectAnswer == "" ||
		question.DifficultyLevel <= 0)
}

func (h *QuestionHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	filters := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&filters)

	questions, err := h.Service.GetQuestions(filters)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch questions")
		return
	}

	response.Success(w, questions)
}

func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var question models.Question

	json.NewDecoder(r.Body).Decode(&question)

	if !ValidateQuestion(question) {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	_, err := h.Service.CreateQuestion(question)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create a question")
		return
	}

	response.Success(w, "Question created successfully")
}

func (h *QuestionHandler) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	var question models.Question

	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.UpdateQuestion(question)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update the question")
		return
	}

	response.Success(w, "Question updated successfully")
}

func (h *QuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	var id int

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.DeleteQuestion(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete the question")
		return
	}

	response.Success(w, "Question deleted successfully")
}
