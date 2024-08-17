package repository_user

import (
	"api-example/internal/domain"
	"context"
)

const CreateUserQuery = `
	INSERT INTO users (id, fullname, email, password, status, created_at, updated_at)
	VALUES (:id, :fullname, :email, :password, :status, :created_at, :updated_at)
`

// CreateUser implements repository.IUserRepository.
func (u *UserRepo) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	model := NewUserDto(user)
	if _, err := u.db.NamedExecContext(ctx, CreateUserQuery, model); err != nil {
		return user, err
	}
	return user, nil
}
