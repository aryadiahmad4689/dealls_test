package mocks

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/repository/user"
	"github.com/stretchr/testify/mock"
)

type UserInterfaceMock struct {
	mock.Mock
}

func Init() user.UserInterface {
	return &UserInterfaceMock{}
}

// GetUserById implements user.UserInterface.
func (m *UserInterfaceMock) GetUserById(ctx context.Context, id string) (entity.User, error) {
	ret := m.Called(ctx, id)
	return ret.Get(0).(entity.User), ret.Error(1)
}
