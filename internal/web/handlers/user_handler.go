package handlers

import (
	"net/http"

	"github.com/Manuelda10/restaurant-backend/internal/application"
)

type UserHandler struct{
	userUseCase application.UserUseCase
}

func NewUserHandler(userUseCase application.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (u *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("RegisterUser"))
}

func (u *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUserByID"))
}

