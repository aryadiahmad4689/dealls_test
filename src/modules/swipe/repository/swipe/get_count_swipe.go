package swipe

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
)

func (s *Swipe) GetCountSwipe(ctx context.Context, req entity.GetSwipeReq) (int, error) {
	var swipeCount int
	formattedDate := req.Date.Format("2006-01-02")

	err := s.db.Slave.QueryRow("SELECT COUNT(*) FROM swipes WHERE swipe_user_id = $1 AND date(created_at) = $2", req.UserId, formattedDate).Scan(&swipeCount)
	if err != nil {
		return 0, err
	}

	return swipeCount, nil
}
