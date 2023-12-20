package endpoint

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/user/entity"
)

func (ed *Endpoint) SignIn(ctx context.Context, req SignInRequest) (interface{}, error) {
	token, err := ed.usecase.SignIn(ctx, entity.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return SignInResp{}, err
	}
	return SignInResp{
		Token: token,
	}, nil
}
