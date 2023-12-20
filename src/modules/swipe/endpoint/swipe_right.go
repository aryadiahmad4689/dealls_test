package endpoint

import (
	"context"
	"time"

	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
)

func (ed *Endpoint) SwipeRight(ctx context.Context, req SwipeRightReq) (interface{}, error) {
	// saya andaikan disini mengambil waktu sekarang, dari local.
	// ini untuk mempersimple saja dan mempercepat development
	now := time.Now()
	user_id := ctx.Value("user_id")
	strValue, _ := user_id.(string)
	err := ed.usecase.SwipeRight(ctx, entity.Swipe{
		SwipeUserId:   strValue,
		IsSwipeUserId: req.IsSwipeUserId,
		CreatedAt:     now,
		SwipeType:     "Like",
	})
	if err != nil {
		return RespNull{}, err
	}
	return RespNull{}, nil
}
