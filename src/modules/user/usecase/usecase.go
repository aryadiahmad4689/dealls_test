package usecase

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/user/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/user/repository"
)

type (
	UseCaseInterface interface {
		SignUp(context.Context, entity.User) (string, error)
		SignIn(ctx context.Context, req entity.User) (string, error)
	}

	UseCase struct {
		app  *app.App
		repo *repository.Repository
	}
)

func Init(app *app.App, repo *repository.Repository) UseCaseInterface {
	return &UseCase{
		app:  app,
		repo: repo,
	}
}
