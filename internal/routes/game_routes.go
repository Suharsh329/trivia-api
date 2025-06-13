package routes

import (
	"database/sql"
	"net/http"
	"trivia/internal/handlers"
	"trivia/internal/services"
)

func RegisterGameRoutes(mux *http.ServeMux, db *sql.DB) {
	gameService := services.NewGameService(db)
	gameHandler := handlers.NewGameHandler(gameService)

	mux.HandleFunc("GET /games", gameHandler.GetGames)
	mux.HandleFunc("GET /game", gameHandler.GetGame)
	mux.HandleFunc("POST /games", gameHandler.CreateGame)
	mux.HandleFunc("POST /games/set-selected-game", gameHandler.SetSelectedGame)
	mux.HandleFunc("POST /games/set-random-game", gameHandler.SetRandomGame)
	mux.HandleFunc("PUT /games", gameHandler.UpdateGame)
	mux.HandleFunc("DELETE /games", gameHandler.DeleteGame)
}
