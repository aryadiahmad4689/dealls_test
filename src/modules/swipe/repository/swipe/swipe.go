package swipe

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
)

type Swipe struct {
	db app.Db
}

type SwipeInterface interface {
	GetSwipeAbleUser(ctx context.Context, req entity.GetSwipeReq) ([]entity.User, error)
	GetCountSwipe(ctx context.Context, req entity.GetSwipeReq) (int, error)
	GetUserSwipe(ctx context.Context, req entity.Swipe) (entity.Swipe, error)
	StoreSwipe(ctx context.Context, req entity.Swipe) (entity.Swipe, error)
}

func Init(app *app.App) SwipeInterface {
	var db = app.GetDb()
	return &Swipe{
		db: *db,
	}
}
