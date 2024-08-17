package repository_user

import (
	"api-example/internal/repository"
	"api-example/pkg/mySqlExt"
)

type UserRepo struct {
	db mySqlExt.IMySqlExt
}

func New(db mySqlExt.IMySqlExt) *UserRepo {
	return &UserRepo{db: db}
}

var _ repository.IUserRepository = (*UserRepo)(nil)
