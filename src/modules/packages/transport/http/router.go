package http

import (
	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/endpoint"
	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/transport/http/handler"
	"github.com/go-chi/chi"
)

func Init(app *app.App, endpoint endpoint.EndpointInterface) *chi.Mux {
	var (
		router = chi.NewRouter()
		h      = handler.InitHandler(app, endpoint)
	)

	router.Get("/", h.GetPackage)

	return router
}
