package routes

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB) {
	RegisterQuestionRoutes(mux, db)
	RegisterGameRoutes(mux, db)
	RegisterCategoryRoutes(mux, db)
	RegisterSubCategoryRoutes(mux, db)
	RegisterPlayerRoutes(mux, db)
	RegisterTeamRoutes(mux, db)
}
