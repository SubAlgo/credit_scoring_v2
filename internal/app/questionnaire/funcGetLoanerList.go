package questionnaire

import (
	"context"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type loanerData struct {
	No            int    `json:"no"`
	LoanerID      int64  `json:"loanerID"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Email         string `json:"email"`
	SendAt        string `json:"sendAt"`
	UpdatedAt     string `json:"updatedAt"`
	UpdatedByID   int64  `json:"updatedByID"`
	UpdatedByName string `json:"updatedByName"`
}

type loanerList struct {
	Total int           `json:"total"`
	List  []*loanerData `json:"list"`
}

type getLoanerListRequest struct {
	statusID int
	name     string
	surname  string
}

func getLoanerList(ctx context.Context, req getLoanerListRequest) (res loanerList, err error) {
	rows, err := dbctx.Query(ctx, `
			SELECT 	users.id, users.name, users.surname, 
					to_char(q.sendAt, 'DD Mon YYYY เวลา HH:MM:SS') as sendAT,
					to_char(q.updatedAt, 'DD Mon YYYY เวลา HH:MM:SS') as updatedAT,
					q.updatedBy as updatedByID,
					(select users.name as updated_by from users where users.id = q.updatedBy) as updatedByName
			FROM users
			LEFT JOIN questionnaire AS q ON users.id = q.loanerID
			WHERE 
				q.statusID = $1 
				AND 
				users.name LIKE $2 || '%'
				AND
				users.surname LIKE $3 || '%'
			ORDER BY updatedAt ASC
		`, req.statusID, req.name, req.surname)

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	err = dbctx.QueryRow(ctx, `
			SELECT 	count(users.id)
			FROM users
			LEFT JOIN questionnaire AS q ON users.id = q.loanerID
			WHERE 
				q.statusID = $1 
				AND 
				users.name LIKE $2 || '%'
				AND
				users.surname LIKE $3 || '%'
		`, req.statusID, req.name, req.surname).Scan(&res.Total)

	if err != nil {
		return res, err
	}
	res.List = make([]*loanerData, 0)
	i := 1

	defer rows.Close()
	for rows.Next() {
		var x loanerData
		err = rows.Scan(
			&x.LoanerID, &x.Name, &x.Surname, &x.SendAt, &x.UpdatedAt, &x.UpdatedByID, &x.UpdatedByName,
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
