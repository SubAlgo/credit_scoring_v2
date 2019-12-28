package questionnaire

import (
	"context"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
)

func questionnaireGetDataByWorker(ctx context.Context, req getDataArgs) (res QuestionnaireStruct, err error) {

	workerID := auth.GetUserID(ctx)
	roleID := auth.GetUserRole(ctx)
	fmt.Println(roleID)
	if workerID == 0 {
		return res, ErrSignInRequired
	}

	switch roleID {
	case 1, 2, 3:
	default:
		return res, ErrPermissionDeny
	}

	if req.LoanerID == 0 {
		fmt.Println("xxx")
		return res, ErrQuestionnaireSelectDataMissingLoanerID
	}

	err = res.getQuestionnaireData(ctx, req.LoanerID)

	if err != nil {
		return res, err
	}

	return
}
