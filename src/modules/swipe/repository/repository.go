package repository

import (
	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/repository/swipe"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/repository/user"
)

type (
	Repository struct {
		Swipe swipe.SwipeInterface
		User  user.UserInterface
	}
)

func Init(app *app.App) *Repository {
	return &Repository{
		Swipe: swipe.Init(app),
		User:  user.Init(app),
	}
}

func (r *Repository) SetSwipe(u user.UserInterface, s swipe.SwipeInterface) *Repository {
	r.User = u
	r.Swipe = s
	return r
}
