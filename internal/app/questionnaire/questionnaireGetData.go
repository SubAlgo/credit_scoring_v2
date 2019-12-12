package questionnaire

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
)

type getDataArgs struct {
	LoanerID int64 `json:"loanerID"`
}

func questionnaireGetData(ctx context.Context, req getDataArgs) (res QuestionnaireStruct, err error) {

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

	if req.LoanerID == 0 {
		return res, ErrQuestionnaireSelectDataMissingLoanerID
	}

	err = res.getQuestionnaireData(ctx, req.LoanerID)

	if err != nil {
		return res, err
	}

	return
}
