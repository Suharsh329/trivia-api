package handlers

import (
	"encoding/json"
	"net/http"
	"trivia/internal/models"
	"trivia/internal/response"
	"trivia/internal/services"
)

type TeamHandler struct {
	Service *services.TeamService
}

func NewTeamHandler(service *services.TeamService) *TeamHandler {
	return &TeamHandler{Service: service}
}

func ValidateTeam(team models.Team) bool {
	return team.Name != ""
}

func (h *TeamHandler) GetTeams(w http.ResponseWriter, r *http.Request) {
	filters := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&filters)

	teams, err := h.Service.GetAllTeams(filters)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch teams")
		return
	}

	response.Success(w, teams)
}

func (h *TeamHandler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	var team models.Team

	json.NewDecoder(r.Body).Decode(&team)

	if !ValidateTeam(team) {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	_, err := h.Service.CreateTeam(team)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create a team")
		return
	}

	response.Success(w, "Team created successfully")
}

func (h *TeamHandler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	var team models.Team

	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.UpdateTeam(team)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update the team")
		return
	}

	response.Success(w, "Team updated successfully")
}

func (h *TeamHandler) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	var id int

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Service.DeleteTeam(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete the team")
		return
	}

	response.Success(w, "Team deleted successfully")
}
