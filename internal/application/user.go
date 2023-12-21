package application

import (
	"context"

	"github.com/Manuelda10/restaurant-backend/internal/domain"
	"github.com/Manuelda10/restaurant-backend/internal/ports"
)

type UserUseCase struct {
	userRepository ports.UserRepository
}

func NewUserUseCase(userRepo ports.UserRepository) *UserUseCase {
	return &UserUseCase{
			userRepository: userRepo,
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user *domain.User) error {
	// Lógica adicional y validaciones
	return uc.userRepository.CreateUser(ctx, user)
}

func (uc *UserUseCase) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	return uc.userRepository.GetUserByID(ctx, id)
}
