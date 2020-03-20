package user

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type getOldForgotPasswordAndAnswerRequest struct {
}

type getOldForgotPasswordAndAnswerResponse struct {
	ForgotPasswordQuestionID int    `json:"forgotPasswordQuestionID"`
	ForgotPasswordAnswer     string `json:"forgotPasswordAnswer"`
}

func getOldForgotPasswordAndAnswer(ctx context.Context, req getOldForgotPasswordAndAnswerRequest) (res getOldForgotPasswordAndAnswerResponse, err error) {
	userID := auth.GetUserID(ctx)
	if userID == 0 {
		return res, ErrUserNotLogin
	}

	err = dbctx.QueryRow(ctx, `
			select forgotPasswordQuestionID, forgotPasswordAns
			from users
			where id = $1
		`, userID).Scan(&res.ForgotPasswordQuestionID, &res.ForgotPasswordAnswer)

	if err != nil {
		return res, ErrSelectForgotPasswordQuestionAndAnswer
	}
	return
}
