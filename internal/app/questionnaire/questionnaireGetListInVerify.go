package questionnaire

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
)

func questionnaireGetListInVerify(ctx context.Context, req getQuestionnaireListRequest) (res loanerList, err error) {

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

	statusID := 3
	res, err = getLoanerList(ctx, statusID)

	if err != nil {
		return res, ErrQuestionnaireGetListNewLoaner
	}

	return res, err
}
