package user

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
)

func (u *User) UpdateUser(ctx context.Context, user entity.User) error {
	query := `UPDATE users SET is_verified = $1 WHERE id = $2`
	_, err := u.db.Master.ExecContext(ctx, query, user.IsVerified, user.Id)
	return err
}
