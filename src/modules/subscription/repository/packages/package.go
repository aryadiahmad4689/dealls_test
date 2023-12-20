package packages

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
)

type Package struct {
	db app.Db
}

type PackageInterface interface {
	GetPackageById(ctx context.Context, req entity.Package) (pkg entity.Package, err error)
}

func Init(app *app.App) PackageInterface {
	var db = app.GetDb()
	return &Package{
		db: *db,
	}
}
