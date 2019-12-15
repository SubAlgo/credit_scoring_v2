package admin

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
)

func getListAdmin(ctx context.Context, req getUserListRequest) (res userList, err error) {

	// check permission access
	{
		workerID := auth.GetUserID(ctx)
		roleID := auth.GetUserRole(ctx)
		if workerID == 0 {
			return res, ErrNotSignIn
		}

		switch roleID {
		case 1, 2, 3:
		default:
			return res, ErrPermissionNotAllow
		}
	}

	role := 2

	res, err = getWorkerList(ctx, role)
	if err != nil {
		return res, ErrGetListEmployee
	}
	return

}
