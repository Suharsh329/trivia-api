package routes

import (
	"database/sql"
	"net/http"
	"trivia/internal/handlers"
	"trivia/internal/services"
)

func RegisterQuestionRoutes(mux *http.ServeMux, db *sql.DB) {
	questionService := services.NewQuestionService(db)
	questionHandler := handlers.NewQuestionHandler(questionService)

	mux.HandleFunc("GET /questions", questionHandler.GetQuestions)
	mux.HandleFunc("POST /questions", questionHandler.CreateQuestion)
	mux.HandleFunc("PUT /questions", questionHandler.UpdateQuestion)
	mux.HandleFunc("DELETE /questions", questionHandler.DeleteQuestion)
}
