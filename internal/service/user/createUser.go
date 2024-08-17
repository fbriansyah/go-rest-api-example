package service_user

import (
	"api-example/internal/domain"
	"context"
)

// CreateUser implements service.IUserService.
func (u *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
