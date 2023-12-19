package ports

import (
	"context"

	"github.com/Manuelda10/restaurant-backend/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, userID int64) (*domain.User, error)
}

type UserUseCase interface {
	RegisterUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, userID int64) (*domain.User, error)
}