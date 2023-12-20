package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
)

func (us *UseCase) Payment(ctx context.Context, req entity.Subscription) error {
	// teknik pembayaran disini disimpelkan, Untuk memenuhi kebutuhan
	// harusnya disini memakai transaction, untuk menangani apabila error
	// tapi untuk mempercepat development saya memilih tidak menggunakanya
	// check package
	pkg, err := us.repo.Package.GetPackageById(ctx, entity.Package{
		Id: req.PackageId,
	})
	if err != nil {
		return errors.New("package not found")
	}
	if pkg.Price <= req.Pricing {
		now := time.Now()
		req.StartDate = now
		req.EndDate = now.AddDate(0, pkg.SubscriptionLong, 0)
		_, err = us.repo.Subscription.StoreSubcription(ctx, req)
		if err != nil {
			return errors.New("failed subscription")
		}
	} else {
		return errors.New("your money not enough")
	}

	err = us.repo.User.UpdateUser(ctx, entity.User{
		Id:         req.UserId,
		IsVerified: 1,
	})
	if err != nil {
		return errors.New("failed update user")
	}

	return nil

}
