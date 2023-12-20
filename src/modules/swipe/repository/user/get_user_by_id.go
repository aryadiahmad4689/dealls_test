package user

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
)

func (u *User) GetUserById(ctx context.Context, id string) (entity.User, error) {
	var user entity.User
	query := `SELECT id, name, age, gender, is_verified
	          FROM users WHERE id = $1`

	err := u.db.Master.QueryRowContext(ctx, query, id).Scan(
		&user.Id, &user.Name, &user.Age, &user.Gender, &user.IsVerified)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
