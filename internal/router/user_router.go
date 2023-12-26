package routes

import (
	"database/sql"

	db "github.com/Manuelda10/restaurant-backend/internal/adapters"
	"github.com/Manuelda10/restaurant-backend/internal/application"
	"github.com/Manuelda10/restaurant-backend/internal/web/handlers"

	"github.com/go-chi/chi/v5"
)

func UserRouter(dbConn *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	userRepo := db.NewUserRepository(dbConn)
	userService := application.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r.Get("/{userID:[0-9]+}", userHandler.GetUserByID)

	return r
}