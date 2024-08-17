package service_user

import (
	"api-example/internal/domain"
	"context"
)

// UpdateUser implements service.IUserService.
func (u *UserService) UpdateUser(ctx context.Context, user *domain.User) error {
	currentUser, err := u.userRepo.GetUserByID(ctx, user.ID)
	if err != nil {
		return err
	}

	// update user data
	currentUser.FullName = user.FullName
	currentUser.Status = user.Status
	if _, err := u.userRepo.UpdateUser(ctx, currentUser); err != nil {
		return err
	}
	return nil
}
