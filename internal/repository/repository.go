package repository

import (
	"api-example/internal/domain"
	"context"

	"github.com/google/uuid"
)

type IUserRepository interface {
	List(ctx context.Context, dest *[]domain.User, queryFilter domain.User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
