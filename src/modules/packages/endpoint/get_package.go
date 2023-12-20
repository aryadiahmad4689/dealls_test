package endpoint

import (
	"context"
)

func (ed *Endpoint) GetPackage(ctx context.Context) (interface{}, error) {
	packages, err := ed.usecase.GetPackage(ctx)
	if err != nil {
		return ResponseNull{}, err
	}
	return GetPackageResponse{
		Packages: packages,
	}, nil
}
