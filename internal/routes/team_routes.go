package routes

import (
	"database/sql"
	"net/http"
	"trivia/internal/handlers"
	"trivia/internal/services"
)

func RegisterTeamRoutes(mux *http.ServeMux, db *sql.DB) {
	teamService := services.NewTeamService(db)
	teamHandler := handlers.NewTeamHandler(teamService)

	mux.HandleFunc("GET /teams", teamHandler.GetTeams)
	mux.HandleFunc("POST /teams", teamHandler.CreateTeam)
	mux.HandleFunc("PUT /teams", teamHandler.UpdateTeam)
	mux.HandleFunc("DELETE /teams", teamHandler.DeleteTeam)
}
