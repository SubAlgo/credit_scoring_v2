package questionnaireOption

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type optionResponse struct {
	AgeOption                         []*option `json:"ageOption"`
	JobOption                         []*option `json:"jobOption"`
	EduOption                         []*option `json:"eduOption"`
	TimeJobOption                     []*option `json:"timeJobOption"`
	FreChangeNameOption               []*option `json:"freChangeNameOption"`
	TimeOfPhoneNumberOption           []*option `json:"timeOfPhoneNumberOption"`
	TimeOfNameInHouseParticularOption []*option `json:"timeOfNameInHouseParticularOption"`
	PayDebtHistoryOption              []*option `json:"payDebtHistoryOption"`
	StatusInHouseParticularOption     []*option `json:"statusInHouseParticularOption"`
	HaveGuarantorOption               []*option `json:"haveGuarantorOption"`
	IamGuarantorOption                []*option `json:"iamGuarantorOption"`
	IncomeTrendOption                 []*option `json:"incomeTrendOption"`
	LoanObjectOption                  []*option `json:"loanObjectOption"`
	ProvinceOption                    []*option `json:"provinceOption"`
}

func getOption(ctx context.Context, req getQuestionnaireOptionRequest) (res optionResponse, err error) {

	// get age option
	{
		rows, err := dbctx.Query(ctx, `select code, title from ageOption`)

		defer rows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.AgeOption = make([]*option, 0)

		for rows.Next() {
			var x option
			err = rows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.AgeOption = append(res.AgeOption, &x)
		}
		rows.Close()
	}

	// get job option
	{
		rows, err := dbctx.Query(ctx, `select code, title from jobOption order by code`)

		defer rows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.JobOption = make([]*option, 0)

		for rows.Next() {
			var x option
			err = rows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.JobOption = append(res.JobOption, &x)
		}
		rows.Close()
	}

	//get eduOption
	{
		rows, err := dbctx.Query(ctx, `select code, title from eduOption order by code`)

		defer rows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.EduOption = make([]*option, 0)

		for rows.Next() {
			var x option
			err = rows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.EduOption = append(res.EduOption, &x)
		}
		rows.Close()
	}

	//get timeJobOption
	{
		rows, err := dbctx.Query(ctx, `select code, title from timeJobOption order by code`)

		defer rows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.TimeJobOption = make([]*option, 0)

		for rows.Next() {
			var x option
			err = rows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.TimeJobOption = append(res.TimeJobOption, &x)
		}
		rows.Close()
	}

	//get freChangeNameOption
	{
		rows, err := dbctx.Query(ctx, `select code, title from freChangeNameOption order by code`)

		defer rows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.FreChangeNameOption = make([]*option, 0)

		for rows.Next() {
			var x option
			err = rows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.FreChangeNameOption = append(res.FreChangeNameOption, &x)
		}
		rows.Close()
	}

	//get timeOfPhoneNumberOption
	{
		rows, err := dbctx.Query(ctx, `select code, title from timeOfPhoneNumberOption order by code`)

		defer rows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.TimeOfPhoneNumberOption = make([]*option, 0)

		for rows.Next() {
			var x option
			err = rows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.TimeOfPhoneNumberOption = append(res.TimeOfPhoneNumberOption, &x)
		}
		rows.Close()
	}

	//get TimeOfNameInHouseParticularOption
	{
		rows, err := dbctx.Query(ctx, `select code, title from TimeOfNameInHouseParticularOption order by code`)

		defer rows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.TimeOfNameInHouseParticularOption = make([]*option, 0)

		for rows.Next() {
			var x option
			err = rows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.TimeOfNameInHouseParticularOption = append(res.TimeOfNameInHouseParticularOption, &x)
		}
		rows.Close()
	}

	//get PayDebtHistoryOption
	{
		rows, err := dbctx.Query(ctx, `select code, title from PayDebtHistoryOption order by code`)

		defer rows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.PayDebtHistoryOption = make([]*option, 0)

		for rows.Next() {
			var x option
			err = rows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.PayDebtHistoryOption = append(res.PayDebtHistoryOption, &x)
		}
		rows.Close()
	}

	//get StatusInHouseParticularOption
	{
		rows, err := dbctx.Query(ctx, `select code, title from StatusInHouseParticularOption order by code`)

		defer rows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.StatusInHouseParticularOption = make([]*option, 0)

		for rows.Next() {
			var x option
			err = rows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.StatusInHouseParticularOption = append(res.StatusInHouseParticularOption, &x)
		}
		rows.Close()
	}

	//get HaveGuarantorOption
	{
		rows, err := dbctx.Query(ctx, `select code, title from HaveGuarantorOption order by code`)

		defer rows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.HaveGuarantorOption = make([]*option, 0)

		for rows.Next() {
			var x option
			err = rows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.HaveGuarantorOption = append(res.HaveGuarantorOption, &x)
		}
		rows.Close()
	}

	//get IamGuarantorOption
	{
		rows, err := dbctx.Query(ctx, `select code, title from IamGuarantorOption order by code`)

		defer rows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.IamGuarantorOption = make([]*option, 0)

		for rows.Next() {
			var x option
			err = rows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.IamGuarantorOption = append(res.IamGuarantorOption, &x)
		}
		rows.Close()
	}

	//get IncomeTrendOption
	{
		rows, err := dbctx.Query(ctx, `select code, title from IncomeTrendOption order by code`)

		defer rows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.IncomeTrendOption = make([]*option, 0)

		for rows.Next() {
			var x option
			err = rows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.IncomeTrendOption = append(res.IncomeTrendOption, &x)
		}
		rows.Close()
	}

	//get LoanObjectOption
	{
		rows, err := dbctx.Query(ctx, `select code, title from LoanObjectOption order by code`)

		defer rows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.LoanObjectOption = make([]*option, 0)

		for rows.Next() {
			var x option
			err = rows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.LoanObjectOption = append(res.LoanObjectOption, &x)
		}
		rows.Close()
	}

	//get ProvinceOption
	{
		rows, err := dbctx.Query(ctx, `select code, title from provinces order by title`)

		defer rows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.ProvinceOption = make([]*option, 0)

		for rows.Next() {
			var x option
			err = rows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.ProvinceOption = append(res.ProvinceOption, &x)
		}
		rows.Close()
	}
	return
}
