package mocks

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/repository/packages"
	"github.com/stretchr/testify/mock"
)

type PackageInterfaceMock struct {
	mock.Mock
}

func Init() packages.PackageInterface {
	return &PackageInterfaceMock{}
}

func (m *PackageInterfaceMock) GetPackage(ctx context.Context) ([]entity.Package, error) {
	ret := m.Called(ctx)
	return ret.Get(0).([]entity.Package), ret.Error(1)
}
