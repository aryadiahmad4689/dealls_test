package modules

import (
	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/interfaces"
	packages "github.com/aryadiahmad4689/dealls_test/src/modules/packages"
	subscription "github.com/aryadiahmad4689/dealls_test/src/modules/subscription"
	swipe "github.com/aryadiahmad4689/dealls_test/src/modules/swipe"
	user "github.com/aryadiahmad4689/dealls_test/src/modules/user"
)

type Modules struct {
	User         interfaces.ModuleInterface
	Swipe        interfaces.ModuleInterface
	Package      interfaces.ModuleInterface
	Subscription interfaces.ModuleInterface
}

func Init(app *app.App) *Modules {
	return &Modules{
		User:         user.Init(app),
		Swipe:        swipe.Init(app),
		Package:      packages.Init(app),
		Subscription: subscription.Init(app),
	}
}
