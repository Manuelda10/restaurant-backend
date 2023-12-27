package routes

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router (dbConn *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	userRoutes := UserRouter(dbConn)
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the API"))
	})
	r.Mount("/users", userRoutes)

	return r
}