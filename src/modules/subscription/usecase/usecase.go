package usecase

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/repository"
)

type (
	UseCaseInterface interface {
		Payment(ctx context.Context, req entity.Subscription) error
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
