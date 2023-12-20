package http

import (
	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/middleware"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/endpoint"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/transport/http/handler"
	"github.com/go-chi/chi"
)

func Init(app *app.App, endpoint endpoint.EndpointInterface) *chi.Mux {
	var (
		router = chi.NewRouter()
		h      = handler.InitHandler(app, endpoint)
	)
	payment := middleware.Auth(h.Payment)
	router.Post("/payment", payment)

	return router
}
