package repository_user

import (
	"api-example/internal/domain"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

const ListUserQuery = `
	SELECT id, fullname, email, password, status, created_at, updated_at, deleted_at 
	FROM users
	WHERE %s
`

// List implements repository.IUserRepository.
func (u *UserRepo) List(ctx context.Context, dest *[]domain.User, queryFilter domain.User) error {
	filter := getQueryFilter(queryFilter)

	usersDto := []UserDto{}

	query := fmt.Sprintf(ListUserQuery, filter)
	slog.Info("List Repo", "query", query)
	err := u.db.SelectContext(ctx, &usersDto, query)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	for _, userDto := range usersDto {
		model := userDto.Domain()
		model.Maskfields()
		*dest = append(*dest, *model)
	}

	return nil
}
