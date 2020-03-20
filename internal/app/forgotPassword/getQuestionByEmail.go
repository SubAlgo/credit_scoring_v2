package forgotPassword

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
	"strings"
)

type getQuestionByEmailRequest struct {
	Email string `json:"email"`
}

type getQuestionByEmailResponse struct {
	Question string `json:"question"`
}

func getQuestionByEmail(ctx context.Context, req getQuestionByEmailRequest) (res getQuestionByEmailResponse, err error) {
	req.Email = strings.TrimSpace(req.Email)
	if req.Email == "" {
		return res, ErrEmailRequired
	}

	err = dbctx.QueryRow(ctx, `
			select f.title
			from users
			left join forgotpasswordquestionoption as f 
				on users.forgotpasswordquestionid = f.id 
			where email = $1
	`, req.Email).Scan(&res.Question)

	if err == sql.ErrNoRows {
		return res, ErrNoResult
	}

	if err != nil {
		fmt.Println(err)
		return res, ErrGetQuestionByEmail
	}
	return
}
