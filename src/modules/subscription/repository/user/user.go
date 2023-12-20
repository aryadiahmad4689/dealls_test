package user

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
)

type User struct {
	db app.Db
}

type UserInterface interface {
	UpdateUser(ctx context.Context, user entity.User) error
}

func Init(app *app.App) UserInterface {
	var db = app.GetDb()
	return &User{
		db: *db,
	}
}
