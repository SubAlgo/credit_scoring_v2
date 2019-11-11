package auth

import (
	"context"
)

type signOutRequest struct {
	Token string `json:"token"`
}

type signOutResponse struct{}

func signOut(ctx context.Context, req signOutRequest) (res signOutResponse, err error) {

	token := ctx.Value(ctxKeyToken{})

	if token == nil {
		return res, ErrTokenRequired
	}
	tokenStr := token.(string)
	err = RedisClient.Del("creditScoring" + tokenStr).Err()

	if err != nil {
		return res, err
	}

	return
}
