package routes

import (
	"database/sql"
	"net/http"
	"trivia/internal/handlers"
	"trivia/internal/services"
)

func RegisterPlayerRoutes(mux *http.ServeMux, db *sql.DB) {
	playerService := services.NewPlayerService(db)
	playerHandler := handlers.NewPlayerHandler(playerService)

	mux.HandleFunc("GET /players", playerHandler.GetPlayers)
	mux.HandleFunc("POST /players", playerHandler.CreatePlayer)
	mux.HandleFunc("PUT /players", playerHandler.UpdatePlayer)
	mux.HandleFunc("DELETE /players", playerHandler.DeletePlayer)
}
