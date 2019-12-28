package questionnaire

import (
	"context"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
)

type getDataArgs struct {
	LoanerID int64 `json:"loanerID"`
}

func questionnaireGetData(ctx context.Context, req getDataArgs) (res QuestionnaireStruct, err error) {
	if req.LoanerID == 0 {
		req.LoanerID = auth.GetUserID(ctx)
	}

	if req.LoanerID == 0 {
		return res, ErrQuestionnaireSelectDataMissingLoanerID
	}

	err = res.getQuestionnaireData(ctx, req.LoanerID)

	if err != nil {
		fmt.Println("err: ", err)
		return res, err
	}

	return
}
