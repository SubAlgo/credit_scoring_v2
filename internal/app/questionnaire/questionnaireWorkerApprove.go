package questionnaire

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
	"strings"
)

//อัพเดต ผลการอนุมัติ
type approveArgs struct {
	LoanerID       int64 `json:"loanerID"`
	Loan           float64
	StatusID       int     `json:"statusID"`
	ApproveComment string  `json:"approveComment"`
	ApproveRate    float64 `json:"approveRate"`
	ApproveTotal   float64 `json:"approveTotal"`
	Interest       float64 `json:"interest"`
	WorkerID       int64
}

func questionnaireWorkerApprove(ctx context.Context, req *approveArgs) (res processResponse, err error) {
	req.WorkerID = auth.GetUserID(ctx)
	roleID := auth.GetUserRole(ctx)

	if req.WorkerID == 0 {
		return res, ErrSignInRequired
	}

	switch roleID {
	case 1:

	default:
		return res, ErrPermissionDeny
	}

	if req.LoanerID == 0 {
		return res, ErrMissingLoanerID
	}

	switch req.StatusID {
	case 5, 6:

	default:
		return res, ErrQuestionnaireStatusIDNotAvailable
	}

	if req.ApproveRate > 100 || req.ApproveRate < 1 {
		return res, ErrApproveRateNotAvailable
	}

	req.ApproveComment = strings.TrimSpace(req.ApproveComment)
	if req.ApproveComment == "" {
		req.ApproveComment = "-"
	}

	if req.ApproveRate <= 0 && req.StatusID == 5 {
		req.ApproveRate = 100
	}

	if req.Interest <= 0 {
		return res, ErrQuestionnaireInterestNotAvailable
	}

	err = dbctx.QueryRow(ctx, `
			select loan from questionnaire where loanerID = $1
	`, req.LoanerID).Scan(&req.Loan)

	if err != nil {
		return res, ErrGetLoan
	}

	req.ApproveTotal = req.Loan * req.ApproveRate / 100

	_, err = dbctx.Exec(ctx, `
		update questionnaire
		set approveBy = $2,
			statusID = $3,
			approveTotal = $4,
			approveRate = $5,
			approveComment = $6,
			interest = $7
		where loanerID = $1 
	`, req.LoanerID, req.WorkerID, req.StatusID, req.ApproveTotal, req.ApproveRate, req.ApproveComment, req.Interest)

	if err != nil {
		return res, ErrQuestionnaireApprove
	}

	res.Message = "บันทึกผลการอนุมัติสินเช่ือสำเร็จ"
	return

}
