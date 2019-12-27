package questionnaire

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type checkQuestionnaireStatusRequest struct {
}

type checkQuestionnaireStatusResponse struct {
	StatusID    int    `json:"statusID"`
	StatusTitle string `json:"statusTitle"`
}

func questionnaireCheckStatus(ctx context.Context, req checkQuestionnaireStatusRequest) (res checkQuestionnaireStatusResponse, err error) {
	userID := auth.GetUserID(ctx)

	if userID == 0 {
		return res, ErrSignInRequired
	}

	err = dbctx.QueryRow(ctx, `
			select q.statusID, qs.title
			from questionnaire as q
			left join questionnaireStatus as qs on q.statusID = qs.id
			where q.loanerID = $1
	`, userID).Scan(&res.StatusID, &res.StatusTitle)

	if err == sql.ErrNoRows {
		res.StatusID = 0
		res.StatusTitle = "ยังไม่ได้ทำแบบสอบถาม"
		return res, nil
	}

	if err != nil {
		fmt.Println(err)
		return res, ErrQuestionnaireGetStatus
	}
	return

}
