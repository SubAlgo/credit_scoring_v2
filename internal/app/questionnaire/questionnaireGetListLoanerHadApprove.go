package questionnaire

import (
	"context"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
)

func questionnaireGetListLoanerHadApprove(ctx context.Context, req getQuestionnaireListRequest) (res loanerList, err error) {

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

	statusID := 5
	getListParam := getLoanerListRequest{statusID: statusID, name: req.Name, surname: req.Surname}
	res, err = getLoanerListDesc(ctx, getListParam)

	if err != nil {
		fmt.Println(err)
		return res, ErrQuestionnaireGetListNewLoaner
	}

	return res, err
}
