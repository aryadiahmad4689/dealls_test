package endpoint

import "context"

func (ed *Endpoint) GetDatings(ctx context.Context) (interface{}, error) {
	users, err := ed.usecase.GetDatings(ctx)
	if err != nil {
		return RespNull{}, err
	}
	return GetDatingResponse{
		Users: users,
	}, nil
}
