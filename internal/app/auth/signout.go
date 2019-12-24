package auth

import "context"

type signOutRequest struct {
	Token string `json:"token"`
}

type signOutResponse struct{}

func signOut(ctx context.Context, req signOutRequest) (res signOutResponse, err error) {
	if req.Token == "" {
		return res, ErrTokenRequired
	}

	err = RedisClient.Del("creditScoring" + req.Token).Err()
	if err != nil {
		return res, err
	}
	return
}
