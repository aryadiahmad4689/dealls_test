package http

import (
	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/middleware"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/endpoint"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/transport/http/handler"
	"github.com/go-chi/chi"
)

func Init(app *app.App, endpoint endpoint.EndpointInterface) *chi.Mux {
	var (
		router = chi.NewRouter()
		h      = handler.InitHandler(app, endpoint)
	)
	authGetDating := middleware.Auth(h.GetDating)
	authSwipeRight := middleware.Auth(h.SwipeRight)
	authSwipeLeft := middleware.Auth(h.SwipeLeft)

	router.Get("/get-dating", authGetDating)
	router.Put("/right/{user_id}", authSwipeRight)
	router.Put("/left/{user_id}", authSwipeLeft)

	return router
}
