package repository_user

import (
	"context"

	"github.com/google/uuid"
)

const DeleteUserQuery = `
	UPDATE users
	SET deleted_at = NOW()
	WHERE id = ?
`

// DeleteUser soft delete user by id
func (u *UserRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if _, err := u.db.ExecContext(ctx, DeleteUserQuery, id); err != nil {
		return err
	}
	return nil
}
