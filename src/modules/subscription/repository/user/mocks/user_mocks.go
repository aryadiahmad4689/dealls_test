package mocks

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/repository/user"
	"github.com/stretchr/testify/mock"
)

type UserInterfaceMock struct {
	mock.Mock
}

func Init() user.UserInterface {
	return &UserInterfaceMock{}
}

func (m *UserInterfaceMock) UpdateUser(ctx context.Context, user entity.User) error {
	ret := m.Called(ctx, user)
	return ret.Error(0)
}
