package user

import (
	"context"
	"database/sql"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
)

type getProfileRequest struct {
	UserID int64 `json:"userID"`
}

func getProfile(ctx context.Context, req getProfileRequest) (res UserStruct, err error) {
	userID := auth.GetUserID(ctx)
	if userID == 0 {
		return res, ErrUserNotLogin
	}

	err = res.getProfile(ctx, userID)

	if err == sql.ErrNoRows {
		return res, ErrGetProfileNoRows
	}

	if err != nil {
		return res, ErrGetProfile
	}

	return
}
