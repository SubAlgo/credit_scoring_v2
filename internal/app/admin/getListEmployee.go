package admin

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
)

type userData struct {
	No      int    `json:"no"`
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Role    string `json:"role"`
}

type userList struct {
	Total int         `json:"total"`
	List  []*userData `json:"list"`
}

type getUserListRequest struct {
}

func getListEmployee(ctx context.Context, req getUserListRequest) (res userList, err error) {

	// check permission access
	{
		workerID := auth.GetUserID(ctx)
		workerRole := auth.GetUserRole(ctx)
		if workerID == 0 {
			return res, ErrNotSignIn
		}

		switch workerRole {
		case 1, 2, 3:
		default:
			return res, ErrPermissionNotAllow
		}
	}

	role := 3

	res, err = getWorkerList(ctx, role)
	if err != nil {
		return res, ErrGetListEmployee
	}
	return

}
