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

func questionnaireGetApproveResult(ctx context.Context, req getApproveResultRequest) (res getApproveResultResponse, err error) {
	userID := auth.GetUserID(ctx)

	if userID == 0 {
		return res, ErrSignInRequired
	}

	err = dbctx.QueryRow(ctx, `
			select loan, approveTotal, interest, approveTotal, loanerPayBack, approveComment from questionnaire where loanerID = $1
	`, userID).Scan(&res.Loan, &res.ApproveTotal, &res.InterestRate, &res.ApproveTotal, &res.Payback, &res.ApproveComment)

	if err != nil {
		return res, ErrQuestionnaireGetApproveResultDB
	}

	return

}
