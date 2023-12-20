package usecase

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/repository"
)

type (
	UseCaseInterface interface {
		GetPackage(ctx context.Context) ([]entity.Package, error)
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
