package endpoint

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/usecase"
)

type (
	EndpointInterface interface {
		GetDatings(ctx context.Context) (interface{}, error)
		SwipeRight(ctx context.Context, req SwipeRightReq) (interface{}, error)
		SwipeLeft(ctx context.Context, req SwipeLeftReq) (interface{}, error)
	}
	Endpoint struct {
		app     *app.App
		usecase usecase.UseCaseInterface
	}

	GetDatingResponse struct {
		Users []entity.User `json:"users"`
	}

	SwipeRightReq struct {
		IsSwipeUserId string
	}

	SwipeLeftReq struct {
		IsSwipeUserId string
	}

	RespNull struct{}
)

func Init(app *app.App, usecase usecase.UseCaseInterface) EndpointInterface {
	return &Endpoint{
		app:     app,
		usecase: usecase,
	}
}
