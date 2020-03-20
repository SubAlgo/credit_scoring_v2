package user

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
	"strings"
)

type changeForgotPasswordQuestionRequest struct {
	QuestionID int    `json:"questionID"`
	Answer     string `json:"answer"`
}

type changeForgotPasswordQuestionResponse struct {
	Message string `json:"message"`
}

func changeForgotPasswordQuestion(ctx context.Context, req changeForgotPasswordQuestionRequest) (res changeForgotPasswordQuestionResponse, err error) {
	userID := auth.GetUserID(ctx)
	if userID == 0 {
		return res, ErrUserNotLogin
	}

	if req.QuestionID == 0 {
		return res, ErrQuestionIDRequired
	}

	req.Answer = strings.TrimSpace(req.Answer)
	if req.Answer == "" {
		return res, ErrForgotPasswordAnswer
	}

	_, err = dbctx.Exec(ctx, `
		update users
		set 
			forgotPasswordQuestionID = $2,
			forgotPasswordAns = $3
		where id = $1
	`, userID, req.QuestionID, req.Answer)

	if err != nil {
		return res, ErrUpdateForgotPasswordQuestion
	}
	res.Message = "แก้ไขข้อมูลสำเร็จ"
	return
}
