package endpoint

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/usecase"
)

type (
	EndpointInterface interface {
		Payment(ctx context.Context, req PaymentRequest) (interface{}, error)
	}
	Endpoint struct {
		app     *app.App
		usecase usecase.UseCaseInterface
	}

	PaymentRequest struct {
		PackageId string `json:"package_id"`
		Ammount   string `json:"ammount"`
	}

	RespNull struct {
	}
)

func Init(app *app.App, usecase usecase.UseCaseInterface) EndpointInterface {
	return &Endpoint{
		app:     app,
		usecase: usecase,
	}
}
