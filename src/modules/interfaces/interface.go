package interfaces

import "github.com/go-chi/chi"

type ModuleInterface interface {
	GetHttpRouter() *chi.Mux
}
