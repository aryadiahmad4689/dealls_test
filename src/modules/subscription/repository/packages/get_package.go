package packages

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
)

func (p *Package) GetPackageById(ctx context.Context, req entity.Package) (pkg entity.Package, err error) {
	err = p.db.Slave.QueryRowContext(ctx, "SELECT id, subscription_type, subscription_long, price FROM packages where id = $1", req.Id).Scan(&pkg.Id, &pkg.SubscriptionType, &pkg.SubscriptionLong, &pkg.Price)
	if err != nil {
		return entity.Package{}, err
	}

	return pkg, nil
}
