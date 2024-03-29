package questionnaire

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
)

func questionnaireGetListWaitApprove(ctx context.Context, req getQuestionnaireListRequest) (res loanerList, err error) {

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

	statusID := 4
	getListParam := getLoanerListRequest{statusID: statusID, name: req.Name, surname: req.Surname}
	res, err = getLoanerListDesc(ctx, getListParam)

	if err != nil {
		return res, ErrQuestionnaireGetListNewLoaner
	}

	return res, err
}
