package swipe

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
)

func (s *Swipe) GetUserSwipe(ctx context.Context, req entity.Swipe) (entity.Swipe, error) {
	swipe := entity.Swipe{}
	formattedDate := req.CreatedAt.Format("2006-01-02")

	err := s.db.Master.QueryRowContext(ctx, "SELECT id, swipe_user_id, is_swipe_user_id, swipe_type FROM swipes WHERE swipe_user_id = $1 AND is_swipe_user_id = $2 AND date(created_at) = $3", req.SwipeUserId, req.IsSwipeUserId, formattedDate).Scan(&swipe.Id, &swipe.SwipeUserId, &swipe.IsSwipeUserId, &swipe.SwipeType)
	if err != nil {
		if err == sql.ErrNoRows {
			return swipe, errors.New("no swipe found")
		}
		return swipe, err
	}

	return swipe, nil
}
