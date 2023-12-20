package endpoint

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/user/usecase"
)

type (
	EndpointInterface interface {
		SignUp(context.Context, SignUpRequest) (interface{}, error)
		SignIn(ctx context.Context, req SignInRequest) (interface{}, error)
	}
	Endpoint struct {
		app     *app.App
		usecase usecase.UseCaseInterface
	}

	SignUpRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Age      string `json:"age"`
		Gender   string `json:"gender"`
	}

	SignUpResp struct {
		Token string `json:"token"`
	}

	SignInRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignInResp struct {
		Token string `json:"token"`
	}
)

func Init(app *app.App, usecase usecase.UseCaseInterface) EndpointInterface {
	return &Endpoint{
		app:     app,
		usecase: usecase,
	}
}
