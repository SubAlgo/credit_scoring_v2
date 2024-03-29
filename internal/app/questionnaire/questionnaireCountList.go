package questionnaire

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type questionnaireCountListRequest struct {
}

type questionnaireCountListResponse struct {
	TotalCustomer          int `json:"totalCustomer"`
	TotalLoanerNotMake     int `json:"totalLoanerNotMake"`
	TotalNewLoaner         int `json:"totalNewLoaner"`
	TotalInVerifyLoaner    int `json:"totalInVerifyLoaner"`
	TotalWaitApproveLoaner int `json:"totalWaitApproveLoaner"`
	TotalLoanerHadApprove  int `json:"totalLoanerHadApprove"`
	TotalLoanerHadDeny     int `json:"totalLoanerHadDeny"`
}

func questionnaireCountList(ctx context.Context, req questionnaireCountListRequest) (res questionnaireCountListResponse, err error) {
	{
		err = dbctx.QueryRow(ctx, `
			select 
				count(id)
			from 
				users 
			where roleID = 4;
		`).Scan(&res.TotalCustomer)

		if err != nil {
			return res, ErrGetTotalCustomer
		}
	}
	{
		err = dbctx.QueryRow(ctx, `
			select 
				count(id)
			from 
				users 
			where id not in (select q.loanerID from questionnaire q where q.statusID > 1) and roleID = 4;
		`).Scan(&res.TotalLoanerNotMake)

		if err != nil {
			return res, ErrGetTotalLoanerNotMake
		}
	}

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

	{
		err = dbctx.QueryRow(ctx, `
			select count(id) from questionnaire where statusID = 5;
		`).Scan(&res.TotalLoanerHadApprove)

		if err != nil {
			return res, ErrGetLoanerHasApprove
		}
	}

	{
		err = dbctx.QueryRow(ctx, `
			select count(id) from questionnaire where statusID = 6;
		`).Scan(&res.TotalLoanerHadDeny)

		if err != nil {
			return res, ErrGetLoanerHasDeny
		}
	}

	return
}
