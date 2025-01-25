package handlers

import (
	"encoding/json"
	"net/http"
	"trivia/internal/models"
	"trivia/internal/response"
	"trivia/internal/services"
)

type GameHandler struct {
	Service *services.GameService
}

func NewGameHandler(service *services.GameService) *GameHandler {
	return &GameHandler{Service: service}
}

func ValidateGame(game models.Game) bool {
	return !(game.Name == "")
}

func (h *GameHandler) GetGames(w http.ResponseWriter, r *http.Request) {
	filters := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&filters)

	games, err := h.Service.GetGames(filters)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch games")
		return
	}

	response.Success(w, games)
}

func (h *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	var game models.Game

	json.NewDecoder(r.Body).Decode(&game)

	if !ValidateGame(game) {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	_, err := h.Service.CreateGame(game)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create a game")
		return
	}

	response.Success(w, "Game created successfully")
}

func (h *GameHandler) UpdateGame(w http.ResponseWriter, r *http.Request) {
	var game models.Game

	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.UpdateGame(game)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update the game")
		return
	}

	response.Success(w, "Game updated successfully")
}

func (h *GameHandler) DeleteGame(w http.ResponseWriter, r *http.Request) {
	var id int

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.DeleteGame(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete the game")
		return
	}

	response.Success(w, "Game deleted successfully")
}

func (h *GameHandler) SetSelectedGame(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		GameId      int   `json:"gameId"`
		QuestionIds []int `json:"questionIds"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.SetSelectedGame(requestData.GameId, requestData.QuestionIds)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to set selected game")
		return
	}

	response.Success(w, "Selected game set successfully")
}

func (h *GameHandler) SetRandomGame(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		GameId            int             `json:"gameId"`
		NumberOfQuestions int             `json:"numberOfQuestions"`
		Percentages       map[int]float64 `json:"percentages"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.SetRandomGame(requestData.GameId, requestData.NumberOfQuestions, requestData.Percentages)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to set random game")
		return
	}

	response.Success(w, "Random game set successfully")
}