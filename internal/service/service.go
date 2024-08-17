package service

import (
	"api-example/internal/domain"
	"context"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
	ListUser(ctx context.Context, user *domain.User) ([]domain.User, error)
}
