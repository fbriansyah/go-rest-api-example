package repository_user

import (
	"api-example/constants"
	"api-example/internal/domain"
	"context"
	"database/sql"
	"log/slog"
)

const (
	getUserByEmailQuery = `SELECT id, fullname, email, password FROM users WHERE email = ?`
)

// GetUserByEmail implements repository.IUserRepository.
func (u *UserRepo) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user UserDto
	err := u.db.GetContext(ctx, &user, getUserByEmailQuery, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, constants.ErrNoRows
		}
		slog.Error("Error when get user by email", "error", err)
		return nil, err
	}

	userDomain := user.Domain()
	return userDomain, nil
}
