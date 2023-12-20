package endpoint

import (
	"context"
	"strconv"

	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
)

func (ed *Endpoint) Payment(ctx context.Context, req PaymentRequest) (interface{}, error) {
	user_id := ctx.Value("user_id")
	strValue, _ := user_id.(string)
	ammount, _ := strconv.ParseFloat(req.Ammount, 64)
	err := ed.usecase.Payment(ctx, entity.Subscription{
		UserId:    strValue,
		PackageId: req.PackageId,
		Pricing:   ammount,
	})
	if err != nil {
		return RespNull{}, err
	}
	return RespNull{}, nil
}
