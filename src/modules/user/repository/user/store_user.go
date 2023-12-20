package user

import (
	"context"
	"time"

	"github.com/aryadiahmad4689/dealls_test/src/modules/user/entity"
)

func (u *User) StoreUser(ctx context.Context, user entity.User) (entity.User, error) {
	query := `INSERT INTO users (email, password, name, age, gender,is_verified, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, email, name, age, gender`

	var now = time.Now().UTC()
	var data = entity.User{}

	err := u.db.Master.QueryRowContext(ctx, query,
		user.Email, user.Password, user.Name, user.Age, user.Gender, user.IsVerified, now, now).
		Scan(&data.Id, &data.Email, &data.Name, &data.Age, &data.Gender)
	if err != nil {
		return entity.User{}, err
	}

	return data, nil
}
