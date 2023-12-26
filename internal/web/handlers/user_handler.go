package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Manuelda10/restaurant-backend/internal/application"
	"github.com/Manuelda10/restaurant-backend/internal/domain"
)

type UserHandler struct {
	userUseCase application.UserUseCase
}

func NewUserHandler(userUseCase *application.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: *userUseCase,
	}
}

//Esto es como un response, request
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *domain.User
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

	user, err := h.userUseCase.GetUserByID(r.Context(), int64(id))
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	json.NewEncoder(w).Encode(user)
}
