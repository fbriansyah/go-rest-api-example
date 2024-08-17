package repository_user

import (
	"api-example/internal/domain"
	"context"
)

const (
	updateUserQuery = `
		UPDATE users
		SET 
			fullname = :fullname,
			status = :status
		WHERE id = :id
	`
)

// UpdateUser implements repository.IUserRepository.
func (u *UserRepo) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	model := NewUserDto(user)
	if _, err := u.db.NamedExecContext(ctx, updateUserQuery, model); err != nil {
		return user, err
	}
	return user, nil
}
