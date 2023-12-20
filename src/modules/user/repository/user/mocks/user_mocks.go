package mocks

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/user/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/user/repository/user"
	"github.com/stretchr/testify/mock"
)

type UserInterfaceMock struct {
	mock.Mock
}

func Init() user.UserInterface {
	return &UserInterfaceMock{}
}

func (m *UserInterfaceMock) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	ret := m.Called(ctx, email)
	return ret.Get(0).(entity.User), ret.Error(1)
}

func (m *UserInterfaceMock) StoreUser(ctx context.Context, user entity.User) (entity.User, error) {
	ret := m.Called(ctx, user)
	return ret.Get(0).(entity.User), ret.Error(1)
}
