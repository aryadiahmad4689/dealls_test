package packages

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/entity"
)

func (p *Package) GetPackage(ctx context.Context) ([]entity.Package, error) {
	rows, err := p.db.Slave.QueryContext(ctx, "SELECT id, subscription_type, subscription_long, price, created_at, updated_at FROM packages")
	if err != nil {
		return []entity.Package{}, err
	}
	// Iterasi melalui hasil query
	var packages []entity.Package
	for rows.Next() {
		var pkg entity.Package
		err = rows.Scan(&pkg.Id, &pkg.SubscriptionType, &pkg.SubscriptionLong, &pkg.Price, &pkg.CreatedAt, &pkg.UpdatedAt)
		if err != nil {
			return []entity.Package{}, err
		}
		packages = append(packages, pkg)
	}

	return packages, nil

}
