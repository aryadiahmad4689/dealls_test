package subscription

import (
	"context"
	"time"

	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
)

func (s *Subscription) StoreSubcription(ctx context.Context, req entity.Subscription) (result entity.Subscription, err error) {
	query := `INSERT INTO subscriptions (user_id, packages_id, StartDate, EndDate, created_at, updated_at) 
          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	now := time.Now()
	// Menjalankan SQL statement
	err = s.db.Master.QueryRowContext(ctx, query,
		req.UserId, req.PackageId, req.StartDate, req.EndDate, now, now).Scan(&result.Id)

	if err != nil {
		return entity.Subscription{}, err
	}

	return result, nil

}
