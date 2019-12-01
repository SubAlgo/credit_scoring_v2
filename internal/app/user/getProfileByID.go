package user

import (
	"context"
	"database/sql"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
)

func getProfileByID(ctx context.Context, req getProfileRequest) (res UserStruct, err error) {

	if req.UserID == 0 {
		return res, ErrUserIdRequired
	}

	userID := auth.GetUserID(ctx)
	if userID == 0 {
		return res, ErrUserNotLogin
	}

	/*
		userRole := auth.GetUserRole(ctx)
		if userRole == 4 {
			return res, ErrNotPermission
		}
	*/

	err = res.getProfile(ctx, req.UserID)

	if err == sql.ErrNoRows {
		return res, ErrGetProfileNoRows
	}

	if err != nil {
		return res, ErrGetProfile
	}

	return
}
