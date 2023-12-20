package endpoint

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/usecase"
)

type (
	EndpointInterface interface {
		GetPackage(ctx context.Context) (interface{}, error)
	}
	Endpoint struct {
		app     *app.App
		usecase usecase.UseCaseInterface
	}

	GetPackageResponse struct {
		Packages []entity.Package `json:"packages"`
	}

	ResponseNull struct{}
)

func Init(app *app.App, usecase usecase.UseCaseInterface) EndpointInterface {
	return &Endpoint{
		app:     app,
		usecase: usecase,
	}
}
