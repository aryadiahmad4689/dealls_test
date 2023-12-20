package repository

import (
	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/repository/packages"
)

type (
	Repository struct {
		Package packages.PackageInterface
	}
)

func Init(app *app.App) *Repository {
	return &Repository{
		Package: packages.Init(app),
	}
}

func (s *Repository) SetPackage(p packages.PackageInterface) *Repository {
	s.Package = p
	return s
}
