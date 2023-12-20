package mocks

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/repository/subscription"
	"github.com/stretchr/testify/mock"
)

type SubscriptionInterfaceMock struct {
	mock.Mock
}

func Init() subscription.SubscriptionInterface {
	return &SubscriptionInterfaceMock{}
}

func (m *SubscriptionInterfaceMock) StoreSubcription(ctx context.Context, req entity.Subscription) (result entity.Subscription, err error) {
	ret := m.Called(ctx, req)
	return ret.Get(0).(entity.Subscription), ret.Error(1)
}
