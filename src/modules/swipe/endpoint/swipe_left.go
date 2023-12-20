package endpoint

import (
	"context"
	"time"

	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
)

func (ed *Endpoint) SwipeLeft(ctx context.Context, req SwipeLeftReq) (interface{}, error) {
	// saya andaikan disini mengambil waktu sekarang, dari local.
	// ini untuk mempersimple saja dan mempercepat development
	now := time.Now()
	user_id := ctx.Value("user_id")
	strValue, _ := user_id.(string)
	err := ed.usecase.SwipeLeft(ctx, entity.Swipe{
		SwipeUserId:   strValue,
		IsSwipeUserId: req.IsSwipeUserId,
		CreatedAt:     now,
		SwipeType:     "Pass",
	})
	if err != nil {
		return RespNull{}, err
	}
	return RespNull{}, nil
}
