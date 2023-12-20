package usecase

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/entity"
)

func (us *UseCase) GetPackage(ctx context.Context) ([]entity.Package, error) {
	packages, err := us.repo.Package.GetPackage(ctx)
	if err != nil {
		return []entity.Package{}, err
	}
	return packages, nil
}
