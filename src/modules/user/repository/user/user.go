package user

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/user/entity"
)

type User struct {
	db app.Db
}

type UserInterface interface {
	StoreUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

func Init(app *app.App) UserInterface {
	var db = app.GetDb()
	return &User{
		db: *db,
	}
}
