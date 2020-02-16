package questionnaire

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type questionnaireCountListRequest struct {
}

type questionnaireCountListResponse struct {
	TotalNewLoaner         int `json:"totalNewLoaner"`
	TotalInVerifyLoaner    int `json:"totalInVerifyLoaner"`
	TotalWaitApproveLoaner int `json:"totalWaitApproveLoaner"`
}

func questionnaireCountList(ctx context.Context, req questionnaireCountListRequest) (res questionnaireCountListResponse, err error) {

	{
		err = dbctx.QueryRow(ctx, `
			select count(id) from questionnaire where statusID = 2;
		`).Scan(&res.TotalNewLoaner)

		if err != nil {
			return res, ErrGetTotalNewLoaner
		}
	}

	{
		err = dbctx.QueryRow(ctx, `
			select count(id) from questionnaire where statusID = 3;
		`).Scan(&res.TotalInVerifyLoaner)

		if err != nil {
			return res, ErrGetTotalVerifyLoaner
		}
	}

	{
		err = dbctx.QueryRow(ctx, `
			select count(id) from questionnaire where statusID = 4;
		`).Scan(&res.TotalWaitApproveLoaner)

		if err != nil {
			return res, ErrGetTotalWaitApproveLoaner
		}
	}
	
	return
}
