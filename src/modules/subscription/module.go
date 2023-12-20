package module

import (
	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/interfaces"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/endpoint"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/repository"
	transporthttp "github.com/aryadiahmad4689/dealls_test/src/modules/subscription/transport/http"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/usecase"
	"github.com/go-chi/chi"
)

type Module struct {
	usecase    usecase.UseCaseInterface
	endpoint   endpoint.EndpointInterface
	repository *repository.Repository
	httpRouter *chi.Mux
}

func Init(app *app.App) interfaces.ModuleInterface {
	var (
		repository = repository.Init(app)
		usecase    = usecase.Init(app, repository)
		endpoint   = endpoint.Init(app, usecase)
		httpRouter = transporthttp.Init(app, endpoint)
	)

	return &Module{
		repository: repository,
		usecase:    usecase,
		endpoint:   endpoint,
		httpRouter: httpRouter,
	}
}

func (module *Module) GetHttpRouter() *chi.Mux {
	return module.httpRouter
}
