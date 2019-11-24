package admin

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type changeUserRoleRequest struct {
	UserID int64 `json:"UserID"`
	RoleID int   `json:"roleID"`
}

func changeUserRole(ctx context.Context, req changeUserRoleRequest) (res processResponse, err error) {
	userRole := auth.GetUserRole(ctx)

	if userRole != 1 {
		return res, ErrPermissionNotAllow
	}

	{
		if req.UserID == 0 || req.RoleID == 0 {
			return res, ErrParamMissing
		}
	}

	_, err = dbctx.Exec(ctx, `
		update users
		set roleID = $2
		where id = $1
	`, req.UserID, req.RoleID)

	if err != nil {
		return res, ErrUpdateUserRole
	}

	res.Message = "เปลี่ยนสถานะสำเร็จ"

	return
}
