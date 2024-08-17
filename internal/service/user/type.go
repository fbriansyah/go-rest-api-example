package service_user

import (
	"api-example/internal/repository"
	"api-example/internal/service"
)

type UserService struct {
	userRepo repository.IUserRepository
}

func New(userRepo repository.IUserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

var _ service.IUserService = (*UserService)(nil)
