package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Manuelda10/restaurant-backend/internal/domain"
)

// UserService es una interfaz que define los métodos que el servicio de usuario debe implementar
type UserUseCase interface {
	CreateUser(ctx context.Context, user domain.User) error
	GetUserByID(ctx context.Context, id int) (domain.User, error)
	// Otros métodos...
}

type UserHandler struct {
	userUseCase UserUseCase
}

func NewUserHandler(userUseCase UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

//Esto es como un response, request
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	if err := h.userUseCase.CreateUser(r.Context(), user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Obtener ID del usuario desde la URL
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
	}

	user, err := h.userUseCase.GetUserByID(r.Context(), id)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	json.NewEncoder(w).Encode(user)
}
