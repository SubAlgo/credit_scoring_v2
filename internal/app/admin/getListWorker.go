package admin

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

func getListAllWorker(ctx context.Context, req getUserListRequest) (res userList, err error) {
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

	rows, err := dbctx.Query(ctx, `
		select users.id, name, surname, email, roles.title
		from users
		left join roles on users.roleID = roles.id
		where roleID = 1 or roleID = 2 or roleID = 3
	`)

	if err != nil {
		return res, err
	}

	err = dbctx.QueryRow(ctx, `
		select count(users.id)
		from users
		left join roles on users.roleID = roles.id
		where roleID = 1 or roleID = 2 or roleID = 3
	`).Scan(&res.Total)

	if err != nil {
		return res, err
	}

	res.List = make([]*userData, 0)
	i := 1

	for rows.Next() {
		var x userData
		err = rows.Scan(
			&x.ID, &x.Name, &x.Surname, &x.Email, &x.Role,
		)
		x.No = i
		i = i + 1

		if err != nil {
			return res, err
		}
		res.List = append(res.List, &x)
	}

	if err := rows.Err(); err != nil {
		return res, err
	}
	return
}
