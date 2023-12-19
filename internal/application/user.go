package application

import (
	"context"

	"github.com/Manuelda10/restaurant-backend/internal/domain"
	"github.com/Manuelda10/restaurant-backend/internal/ports"
)

type UserUseCase struct{
	userRepository ports.UserRepository
}

func NewUserUseCase(userRepository ports.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (u *UserUseCase) RegisterUser(ctx context.Context, user *domain.User) error {
	return u.userRepository.CreateUser(ctx, user)
}

func (u *UserUseCase) GetUserByID(ctx context.Context, userID int64) (*domain.User, error) {
	return u.userRepository.GetUserByID(ctx, userID)
}