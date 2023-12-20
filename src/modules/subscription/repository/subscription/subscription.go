package subscription

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
)

type Subscription struct {
	db app.Db
}

type SubscriptionInterface interface {
	StoreSubcription(ctx context.Context, req entity.Subscription) (result entity.Subscription, err error)
}

func Init(app *app.App) SubscriptionInterface {
	var db = app.GetDb()
	return &Subscription{
		db: *db,
	}
}
