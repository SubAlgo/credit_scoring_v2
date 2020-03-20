package forgotPassword

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

func getForgotPasswordQuestionOption(ctx context.Context, req forgotPasswordRequest) (res forgotPasswordQuestionList, err error) {

	rows, err := dbctx.Query(ctx, `
			select 	id, title
			from forgotPasswordQuestionOption
		`)

	if err != nil {
		return res, err
	}

	res.List = make([]*forgotPasswordQuestion, 0)

	defer rows.Close()
	for rows.Next() {
		var x forgotPasswordQuestion
		err = rows.Scan(
			&x.ID, &x.Question,
		)

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
