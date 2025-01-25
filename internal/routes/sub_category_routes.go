package routes

import (
	"database/sql"
	"net/http"
	"trivia/internal/handlers"
	"trivia/internal/services"
)

func RegisterSubCategoryRoutes(mux *http.ServeMux, db *sql.DB) {
	subCategoryService := services.NewSubCategoryService(db)
	subCategoryHandler := handlers.NewSubCategoryHandler(subCategoryService)

	mux.HandleFunc("GET /sub_categories", subCategoryHandler.GetSubCategories)
	mux.HandleFunc("POST /sub_categories", subCategoryHandler.CreateSubCategory)
	mux.HandleFunc("PUT /sub_categories", subCategoryHandler.UpdateSubCategory)
	mux.HandleFunc("DELETE /sub_categories", subCategoryHandler.DeleteSubCategory)
}
