package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Manuelda10/restaurant-backend/internal/domain"
)

// UserRepositoryImpl implementa la interfaz ports.UserRepository
type UserRepositoryImpl struct {
    db *sql.DB
}

// NewUserRepository crea una nueva instancia de UserRepositoryImpl
func NewUserRepository(db *sql.DB) *UserRepositoryImpl{
    return &UserRepositoryImpl{db: db}
}

// CreateUser inserta un nuevo usuario en la base de datos
func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *domain.User) error {
    query := `INSERT INTO users (username, password, email, phone_number, created_on, created_by) VALUES (?, ?, ?, ?, ?, ?)`
    _, err := r.db.ExecContext(ctx, query, user.Username, user.Password, user.Email, user.PhoneNumber, user.CreatedOn, user.CreatedBy)
    return err
}

// GetUserByID recupera un usuario por su ID de la base de datos
func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
    query := `SELECT id, username, password, email, phone_number, created_on, created_by, updated_on, updated_by FROM users WHERE id = ?`
    var user *domain.User
    err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.PhoneNumber, &user.CreatedOn, &user.CreatedBy, &user.UpdatedOn, &user.UpdatedBy)
    if err != nil {
        if err == sql.ErrNoRows {
            return &domain.User{}, errors.New("user not found")
        }
        return &domain.User{}, err
    }
    return user, nil
}

// UpdateUser implementa UserRepository.UpdateUser
func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, user *domain.User) error {
    // Implementa la lógica para actualizar un usuario
    // ...
    return nil
}

// DeleteUser implementa UserRepository.DeleteUser
func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, userID int64) error {
    // Implementa la lógica para eliminar un usuario
    // ...
    return nil
}

// Aquí puedes agregar las implementaciones de otros métodos de la interfaz, como UpdateUser y DeleteUser
