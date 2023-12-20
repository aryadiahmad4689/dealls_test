package handler

import (
	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/endpoint"
)

type Handler struct {
	app      *app.App
	endpoint endpoint.EndpointInterface
}

func InitHandler(app *app.App, endpoint endpoint.EndpointInterface) *Handler {
	var handler = &Handler{
		app:      app,
		endpoint: endpoint,
	}
	return handler
}
