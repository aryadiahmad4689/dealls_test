package repository

import (
	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/repository/packages"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/repository/subscription"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/repository/user"
)

type (
	Repository struct {
		Subscription subscription.SubscriptionInterface
		User         user.UserInterface
		Package      packages.PackageInterface
	}
)

func Init(app *app.App) *Repository {
	return &Repository{
		Subscription: subscription.Init(app),
		User:         user.Init(app),
		Package:      packages.Init(app),
	}
}

func (r *Repository) SetSubsription(s subscription.SubscriptionInterface) *Repository {
	r.Subscription = s
	return r
}
func (r *Repository) SetPackage(p packages.PackageInterface) *Repository {
	r.Package = p
	return r
}

func (r *Repository) SetUser(u user.UserInterface) *Repository {
	r.User = u
	return r
}
