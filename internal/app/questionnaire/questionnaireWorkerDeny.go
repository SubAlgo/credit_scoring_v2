package questionnaire

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"strings"
)

func questionnaireWorkerDeny(ctx context.Context, req *approveArgs) (res processResponse, err error) {

	req.QuestionnaireStatusID = 6

	// check signIn
	{
		req.WorkerID = auth.GetUserID(ctx)
		roleID := auth.GetUserRole(ctx)

		if req.WorkerID == 0 {
			return res, ErrSignInRequired
		}

		if roleID != 1 {
			return res, ErrPermissionDeny
		}
	}

	// check input data
	{
		if req.LoanerID == 0 {
			return res, ErrMissingLoanerID
		}
	}

	req.ApproveRate = 0
	req.InterestRate = 0

	req.ApproveComment = strings.TrimSpace(req.ApproveComment)
	if req.ApproveComment == "" {
		req.ApproveComment = "-"
	}

	req.ApproveTotal = 0

	req.LoanerPayback = 0

	err = req.setApproveResult(ctx)

	if err != nil {
		return res, ErrQuestionnaireApprove
	}

	res.Message = "บันทึกผลการอนุมัติสินเช่ือสำเร็จ"

	return

}
