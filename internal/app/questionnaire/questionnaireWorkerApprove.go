package questionnaire

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
	"strings"
)

//อัพเดต ผลการอนุมัติ
type approveArgs struct {
	LoanerID       int64   `json:"loanerID"`
	Loan           float64 `json:"loan"`
	ApproveRate    float64 `json:"approveRate"`
	ApproveTotal   float64
	InterestRate   float64 `json:"interestRate"`
	LoanerPayback  float64
	ApproveComment string `json:"approveComment"`
}

func questionnaireWorkerApprove(ctx context.Context, req *approveArgs) (res processResponse, err error) {

	questionnaireStatusID := 5

	// check signIn
	var workerID int64
	{
		workerID = auth.GetUserID(ctx)
		roleID := auth.GetUserRole(ctx)

		if workerID == 0 {
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

		if req.ApproveRate > 100 || req.ApproveRate < 1 {
			return res, ErrApproveRateNotAvailable
		}

		if req.InterestRate < 0 {
			return res, ErrQuestionnaireInterestNotAvailable
		}
	}

	req.ApproveComment = strings.TrimSpace(req.ApproveComment)
	if req.ApproveComment == "" {
		req.ApproveComment = "-"
	}
	var loan float64
	err = dbctx.QueryRow(ctx, `
				select loanW from questionnaire where loanerID = $1
	`, req.LoanerID).Scan(&loan)

	if err != nil {
		return res, ErrGetLoan
	}

	if req.Loan != loan {
		return res, ErrLoanMustBeNumber
	}

	req.ApproveTotal = req.Loan * (req.ApproveRate / 100)

	req.LoanerPayback = req.ApproveTotal + (req.ApproveTotal * (req.InterestRate / 100))

	_, err = dbctx.Exec(ctx, `
			update questionnaire
			set approveBy = $2,
				statusID = $3,
				approveRate = $4,
				approveTotal = $5,
				approveComment = $6,
				interest = $7,
				loanerPayback = $8
			where loanerID = $1
		`, req.LoanerID, workerID, questionnaireStatusID, req.ApproveRate, req.ApproveTotal, req.ApproveComment, req.InterestRate, req.LoanerPayback)

	if err != nil {
		return res, ErrQuestionnaireApprove
	}

	res.Message = "บันทึกผลการอนุมัติสินเช่ือสำเร็จ"

	return

}
