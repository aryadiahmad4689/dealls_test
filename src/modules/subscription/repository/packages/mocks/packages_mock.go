package mocks

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/repository/packages"
	"github.com/stretchr/testify/mock"
)

type PackageInterfaceMock struct {
	mock.Mock
}

func Init() packages.PackageInterface {
	return &PackageInterfaceMock{}
}

func (m *PackageInterfaceMock) GetPackageById(ctx context.Context, req entity.Package) (pkg entity.Package, err error) {
	ret := m.Called(ctx, req)
	return ret.Get(0).(entity.Package), ret.Error(1)
}
