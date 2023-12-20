package user

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
)

type User struct {
	db app.Db
}

type UserInterface interface {
	GetUserById(ctx context.Context, id string) (entity.User, error)
}

func Init(app *app.App) UserInterface {
	var db = app.GetDb()
	return &User{
		db: *db,
	}
}
