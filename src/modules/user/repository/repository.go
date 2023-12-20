package repository

import (
	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/user/repository/user"
)

type (
	Repository struct {
		app  *app.App
		User user.UserInterface
	}
)

func Init(app *app.App) *Repository {
	return &Repository{
		app:  app,
		User: user.Init(app),
	}
}

func (r *Repository) SetUser(u user.UserInterface) *Repository {
	r.User = u
	return r
}
