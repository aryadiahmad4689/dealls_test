package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
)

func (us *UseCase) GetDatings(ctx context.Context) ([]entity.User, error) {
	// saya andaikan disini mengambil waktu sekarang, dari local.
	// ini untuk mempersimple saja dan mempercepat development
	now := time.Now()
	user_id := ctx.Value("user_id")
	strValue, _ := user_id.(string)
	count, err := us.repo.Swipe.GetCountSwipe(ctx, entity.GetSwipeReq{
		UserId: strValue,
		Date:   now,
	})
	if err != nil {
		return []entity.User{}, err
	}
	// check user
	data, err := us.repo.User.GetUserById(ctx, strValue)
	if err != nil {
		return []entity.User{}, err
	}
	if data.IsVerified == 0 {
		if count >= 10 {
			return []entity.User{}, errors.New("swipe today is done")
		}
	}

	user, err := us.repo.Swipe.GetSwipeAbleUser(ctx, entity.GetSwipeReq{
		UserId: strValue,
		Date:   now,
	})

	if err != nil {
		return []entity.User{}, err
	}

	return user, nil
}
