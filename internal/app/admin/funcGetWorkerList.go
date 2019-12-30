package admin

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

func getWorkerList(ctx context.Context, roleID int) (res userList, err error) {

	rows, err := dbctx.Query(ctx, `
		select users.id, name, surname, email, phone, roles.title
		from users
		left join roles on users.roleID = roles.id
		where users.roleID = $1
	`, roleID)

	if err != nil {
		return res, err
	}

	err = dbctx.QueryRow(ctx, `
		select count(users.id)
		from users
		left join roles on users.roleID = roles.id
		where users.roleID = $1
	`, roleID).Scan(&res.Total)

	if err != nil {
		return res, err
	}

	res.List = make([]*userData, 0)
	i := 1

	for rows.Next() {
		var x userData
		err = rows.Scan(
			&x.ID, &x.Name, &x.Surname, &x.Email, &x.Phone, &x.Role,
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
