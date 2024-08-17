package repository_user

import (
	"api-example/constants"
	"api-example/internal/domain"
	"context"
	"database/sql"
	"log/slog"

	"github.com/google/uuid"
)

const (
	getUserByIDQuery = `SELECT id, fullname, email, password FROM users WHERE id = ?`
)

// GetUserByID implements repository.IUserRepository.
func (u *UserRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var user UserDto
	err := u.db.GetContext(ctx, &user, getUserByIDQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, constants.ErrNoRows
		}
		slog.Error("Error when get user by id", "error", err)
		return nil, err
	}

	userDomain := user.Domain()
	return userDomain, nil
}
