package questionnaire

import (
	"context"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

func questionnaireGetListNewLoanerNotMakeQuestionnaire(ctx context.Context, req getQuestionnaireListRequest) (res loanerList, err error) {

	// check permission access
	{
		workerID := auth.GetUserID(ctx)
		roleID := auth.GetUserRole(ctx)

		if workerID == 0 {
			return res, ErrSignInRequired
		}

		switch roleID {
		case 1, 2, 3:

		default:
			return res, ErrPermissionDeny
		}
	}

	rows, err := dbctx.Query(ctx, `
			SELECT 
				id, name, surname, email
			FROM 
				users 
			WHERE 
				id NOT IN (select q.loanerID from questionnaire q where q.statusID > 1) 
				AND 
					roleID = 4
				AND 
					name LIKE $1 || '%'
				AND 
					surname LIKE $2 || '%'
		`, req.Name, req.Surname)
	if err != nil {
		return res, err
	}

	err = dbctx.QueryRow(ctx, `
			SELECT 
				count(id)
			FROM 
				users 
			WHERE 
				id NOT IN (select q.loanerID from questionnaire q where q.statusID > 1) 
				AND 
				roleID = 4
			AND 
				name LIKE $1 || '%'
			AND 
				surname LIKE $2 || '%'
		`, req.Name, req.Surname).Scan(&res.Total)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	res.List = make([]*loanerData, 0)
	i := 1

	defer rows.Close()
	for rows.Next() {
		var x loanerData
		err = rows.Scan(
			&x.LoanerID, &x.Name, &x.Surname, &x.Email,
		)

		x.No = i
		i = i + 1
		if err != nil {
			return res, err
		}
		res.List = append(res.List, &x)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return res, err
	}
	return
}
