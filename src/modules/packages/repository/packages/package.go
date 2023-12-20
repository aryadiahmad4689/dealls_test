package packages

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/entity"
)

type Package struct {
	db app.Db
}

type PackageInterface interface {
	GetPackage(ctx context.Context) ([]entity.Package, error)
}

func Init(app *app.App) PackageInterface {
	var db = app.GetDb()
	return &Package{
		db: *db,
	}
}
