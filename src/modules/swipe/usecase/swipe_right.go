package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
)

func (us *UseCase) SwipeRight(ctx context.Context, req entity.Swipe) error {
	// check swipe apakah sudah pernah dilakukan kepada user yang sama
	// dan di waktu yang sama. anggap saja ini langkah pencegahan, walaupun
	// kita sudah memastikan bahwa tidak ada user yang sama yg sama yang akan di tampilkan
	// sebanyak dua kali.
	req.CreatedAt = time.Now()
	_, err := us.repo.Swipe.GetUserSwipe(ctx, req)
	if err == nil {
		return errors.New("already swipped")
	}
	_, err = us.repo.Swipe.StoreSwipe(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
