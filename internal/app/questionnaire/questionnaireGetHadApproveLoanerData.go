package questionnaire

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
)

func questionnaireGetHadApproveLoanerData(ctx context.Context, req getDataArgs) (res QuestionnaireStruct, err error) {

	roleID := auth.GetUserRole(ctx)
	if roleID == 4 {
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
