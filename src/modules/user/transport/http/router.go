package http

import (
	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/user/endpoint"
	"github.com/aryadiahmad4689/dealls_test/src/modules/user/transport/http/handler"
	"github.com/go-chi/chi"
)

func Init(app *app.App, endpoint endpoint.EndpointInterface) *chi.Mux {
	var (
		router = chi.NewRouter()
		h      = handler.InitHandler(app, endpoint)
	)

	router.Post("/sign-up", h.SignUp)
	router.Post("/sign-in", h.SignIn)

	return router
}
