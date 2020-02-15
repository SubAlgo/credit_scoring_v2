package questionnaire

import (
	"context"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type sendBackToVerifyRequest struct {
	LoanerID int64 `json:"loanerID"`
}

type sendBackToVerifyResponse struct {
	Message string `json:"message"`
}

func sendBackToVerify(ctx context.Context, req sendBackToVerifyRequest) (res sendBackToVerifyResponse, err error) {
	workerID := auth.GetUserID(ctx)
	roleID := auth.GetUserRole(ctx)

	if workerID == 0 {
		return res, ErrSignInRequired
	}

	if roleID != 1 {
		return res, ErrPermissionDeny
	}

	if req.LoanerID == 0 {
		return res, ErrSendBackToVerifyLoanerIDRequest
	}

	_, err = dbctx.Exec(ctx, `
				update questionnaire set statusID = 3 where loanerID = $1;
			`, req.LoanerID)

	if err != nil {
		fmt.Println(err)
		return res, ErrSendBackToVerifyUpdateDB
	}
	res.Message = "ทำรายการสำเร็จ"
	return
}
