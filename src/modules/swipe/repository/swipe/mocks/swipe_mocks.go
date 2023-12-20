package mocks

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/repository/swipe"
	"github.com/stretchr/testify/mock"
)

type SwipeInterfaceMock struct {
	mock.Mock
}

func Init() swipe.SwipeInterface {
	return &SwipeInterfaceMock{}
}

// GetCountSwipe implements swipe.SwipeInterface.
func (m *SwipeInterfaceMock) GetCountSwipe(ctx context.Context, req entity.GetSwipeReq) (int, error) {
	ret := m.Called(ctx, req)
	return ret.Get(0).(int), ret.Error(1)
}

// GetSwipeAbleUser implements swipe.SwipeInterface.
func (m *SwipeInterfaceMock) GetSwipeAbleUser(ctx context.Context, req entity.GetSwipeReq) ([]entity.User, error) {
	ret := m.Called(ctx, req)
	return ret.Get(0).([]entity.User), ret.Error(1)
}

// GetUserSwipe implements swipe.SwipeInterface.
func (m *SwipeInterfaceMock) GetUserSwipe(ctx context.Context, req entity.Swipe) (entity.Swipe, error) {
	ret := m.Called(ctx, req)
	return ret.Get(0).(entity.Swipe), ret.Error(1)
}

// StoreSwipe implements swipe.SwipeInterface.
func (m *SwipeInterfaceMock) StoreSwipe(ctx context.Context, req entity.Swipe) (entity.Swipe, error) {
	ret := m.Called(ctx, req)
	return ret.Get(0).(entity.Swipe), ret.Error(1)
}
