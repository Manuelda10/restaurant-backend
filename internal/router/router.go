package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Router (dbConn *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	userRoutes := UserRouter(dbConn)
	
	r.Mount("/users", userRoutes)

	return r
}