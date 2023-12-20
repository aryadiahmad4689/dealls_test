package endpoint

import (
	"context"
	"strconv"

	"github.com/aryadiahmad4689/dealls_test/src/modules/user/entity"
)

func (ed *Endpoint) SignUp(ctx context.Context, req SignUpRequest) (interface{}, error) {
	age, _ := strconv.Atoi(req.Age)
	token, err := ed.usecase.SignUp(ctx, entity.User{
		Name:       req.Name,
		Email:      req.Email,
		Password:   req.Password,
		Age:        age,
		Gender:     req.Gender,
		IsVerified: 0,
	})
	if err != nil {
		return SignUpResp{}, err
	}
	return SignUpResp{
		Token: token,
	}, nil
}
