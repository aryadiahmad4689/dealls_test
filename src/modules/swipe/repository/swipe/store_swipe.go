package swipe

import (
	"context"
	"time"

	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
)

func (s *Swipe) StoreSwipe(ctx context.Context, req entity.Swipe) (entity.Swipe, error) {
	query := `INSERT INTO swipes (swipe_user_id, is_swipe_user_id, swipe_type,created_at,updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	now := time.Now()
	var newSwipe entity.Swipe
	// Menjalankan SQL statement
	err := s.db.Master.QueryRowContext(ctx, query, req.SwipeUserId, req.IsSwipeUserId, req.SwipeType, now, now).Scan(&newSwipe.Id)
	if err != nil {
		return entity.Swipe{}, err
	}

	return newSwipe, nil
}
