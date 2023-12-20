package usecase

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/repository"
)

type (
	UseCaseInterface interface {
		GetDatings(ctx context.Context) ([]entity.User, error)
		SwipeRight(ctx context.Context, req entity.Swipe) error
		SwipeLeft(ctx context.Context, req entity.Swipe) error
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
