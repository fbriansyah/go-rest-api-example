package repository_user

import (
	"api-example/internal/domain"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// UserDto is a struct that represents a user in the database.
type UserDto struct {
	ID        uuid.UUID    `db:"id"`
	FullName  string       `db:"fullname"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	Status    string       `db:"status"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

// NewUserDto creates a new UserDto from a domain.User
func NewUserDto(user *domain.User) UserDto {
	deletedAt := sql.NullTime{
		Valid: false,
	}
	if user.DeletedAt != nil {
		deletedAt = sql.NullTime{
			Valid: true,
			Time:  *user.DeletedAt,
		}
	}
	return UserDto{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  user.Password,
		Status:    string(user.Status),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: deletedAt,
	}
}

// Domain converts a UserDto to a domain.User
func (u *UserDto) Domain() *domain.User {
	var deletedAt *time.Time
	if u.DeletedAt.Valid {
		deletedAt = &u.DeletedAt.Time
	}
	return &domain.User{
		ID:        u.ID,
		FullName:  u.FullName,
		Email:     u.Email,
		Password:  u.Password,
		Status:    domain.UserStatus(u.Status),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: deletedAt,
	}
}
