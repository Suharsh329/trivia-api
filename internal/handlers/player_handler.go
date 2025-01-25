package handlers

import (
	"encoding/json"
	"net/http"
	"trivia/internal/models"
	"trivia/internal/response"
	"trivia/internal/services"
)

type PlayerHandler struct {
	Service *services.PlayerService
}

func NewPlayerHandler(service *services.PlayerService) *PlayerHandler {
	return &PlayerHandler{Service: service}
}

func ValidatePlayer(player models.Player) bool {
	return player.Name != ""
}

func (h *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	filters := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&filters)

	players, err := h.Service.GetAllPlayers(filters)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch players")
		return
	}

	response.Success(w, players)
}

func (h *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player

	json.NewDecoder(r.Body).Decode(&player)

	if !ValidatePlayer(player) {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	_, err := h.Service.CreatePlayer(player)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create a player")
		return
	}

	response.Success(w, "Player created successfully")
}

func (h *PlayerHandler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player

	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.UpdatePlayer(player)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update the player")
		return
	}

	response.Success(w, "Player updated successfully")
}

func (h *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	var id int

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.DeletePlayer(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete the player")
		return
	}

	response.Success(w, "Player deleted successfully")
}
