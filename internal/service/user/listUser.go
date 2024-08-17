package service_user

import (
	"api-example/internal/domain"
	"context"
)

// ListUser implements service.IUserService.
func (u *UserService) ListUser(ctx context.Context, user *domain.User) ([]domain.User, error) {
	users := []domain.User{}
	err := u.userRepo.List(ctx, &users, *user)
	if err != nil {
		return users, err
	}

	return users, nil
}
