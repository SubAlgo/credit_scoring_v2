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
			select 
				id, name, surname, email
			from 
				users 
			where id not in (select q.loanerID from questionnaire q) and roleID = 4;
		`)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	err = dbctx.QueryRow(ctx, `
			select 
				count(id)
			from 
				users 
			where id not in (select q.loanerID from questionnaire q) and roleID = 4;
		`).Scan(&res.Total)
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
