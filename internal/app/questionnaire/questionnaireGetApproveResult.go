package questionnaire

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type getApproveResultRequest struct {
}

type getApproveResultResponse struct {
	Loan           float64 `json:"loan"`
	ApproveTotal   float64 `json:"approveTotal"`
	InterestRate   float64 `json:"interestRate"`
	Payback        float64 `json:"payback"`
	ApproveComment string  `json:"approveComment"`
}

type nullHandle struct {
	loan           NullFloat64
	interest       NullFloat64
	approveTotal   NullFloat64
	loanerPayBack  NullFloat64
	approveComment NullString
}

func questionnaireGetApproveResult(ctx context.Context, req getApproveResultRequest) (res getApproveResultResponse, err error) {
	userID := auth.GetUserID(ctx)

	if userID == 0 {
		return res, ErrSignInRequired
	}
	var d nullHandle
	err = dbctx.QueryRow(ctx, `
			select loan, approveTotal, interest, loanerPayBack, approveComment from questionnaire where loanerID = $1
	`, userID).Scan(&d.loan, &d.approveTotal, &d.interest, &d.loanerPayBack, &d.approveComment)

	res.Loan = d.loan.Float64
	res.ApproveTotal = d.approveTotal.Float64
	res.InterestRate = d.interest.Float64
	res.Payback = d.loanerPayBack.Float64
	res.ApproveComment = d.approveComment.String

	if err != nil {
		return res, ErrQuestionnaireGetApproveResultDB
	}

	return

}
