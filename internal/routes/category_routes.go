package routes

import (
	"database/sql"
	"net/http"
	"trivia/internal/handlers"
	"trivia/internal/services"
)

func RegisterCategoryRoutes(mux *http.ServeMux, db *sql.DB) {
	categoryService := services.NewCategoryService(db)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	mux.HandleFunc("GET /categories", categoryHandler.GetCategories)
	mux.HandleFunc("GET /categories/{id}", categoryHandler.GetCategoryById)
	mux.HandleFunc("POST /categories", categoryHandler.CreateCategory)
	mux.HandleFunc("PUT /categories", categoryHandler.UpdateCategory)
	mux.HandleFunc("DELETE /categories", categoryHandler.DeleteCategory)
}
