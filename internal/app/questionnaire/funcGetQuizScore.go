package questionnaire

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

func getAgeScore(ctx context.Context, ageCode string) (ageScore int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from ageOption where code = $1
	`, ageCode).Scan(&ageScore)
	return
}

func getJobScore(ctx context.Context, jobCode string) (jobScore int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from jobOption where code = $1
	`, jobCode).Scan(&jobScore)
	return
}

func getEduScore(ctx context.Context, eduCode string) (eduScore int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from eduOption where code = $1
	`, eduCode).Scan(&eduScore)
	return
}

func getTimeJobScore(ctx context.Context, timeJobCode string) (timeJobScore int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from timeJobOption where code = $1
	`, timeJobCode).Scan(&timeJobScore)
	return
}

func getFreChangeNameScore(ctx context.Context, code string) (score int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from freChangeNameOption where code = $1
	`, code).Scan(&score)
	return
}

func getTimeOfPhoneNumberScore(ctx context.Context, code string) (score int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from timeOfPhoneNumberOption where code = $1
	`, code).Scan(&score)
	return
}

func getTimeOfNameInHouseParticularScore(ctx context.Context, code string) (score int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from timeOfNameInHouseParticularOption where code = $1
	`, code).Scan(&score)
	return
}

func getPayDebtHistoryScore(ctx context.Context, code string) (score int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from payDebtHistoryOption where code = $1
	`, code).Scan(&score)
	return
}

func getStatusInHouseParticularScore(ctx context.Context, code string) (score int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from statusInHouseParticularOption where code = $1
	`, code).Scan(&score)
	return
}

func getIncomePerDebtScore(ctx context.Context, code string) (score int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from incomePerDebtOption where code = $1
	`, code).Scan(&score)
	return
}

func getTotalDebtPerYearIncomeScore(ctx context.Context, code string) (score int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from totalDebtPerYearIncomeOption where code = $1
	`, code).Scan(&score)
	return
}

func getSavingPerLoanScore(ctx context.Context, code string) (score int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from savingPerLoanOption where code = $1
	`, code).Scan(&score)
	return
}

func getMortgageSecuritiesPerLoanScore(ctx context.Context, code string) (score int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from mortgageSecuritiesPerLoanOption where code = $1
	`, code).Scan(&score)
	return
}

func getHaveGuarantorScore(ctx context.Context, code string) (score int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from haveGuarantorOption where code = $1
	`, code).Scan(&score)
	return
}

func getIamGuarantorScore(ctx context.Context, code string) (score int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from iamGuarantorOption where code = $1
	`, code).Scan(&score)
	return
}

func getIncomeTrendScore(ctx context.Context, code string) (score int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from incomeTrendOption where code = $1
	`, code).Scan(&score)
	return
}

func getLoanObjectScore(ctx context.Context, code string) (score int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from loanObjectOption where code = $1
	`, code).Scan(&score)
	return
}

func getProvinceScore(ctx context.Context, code string) (score int, err error) {
	err = dbctx.QueryRow(ctx, `
		select score from provinces where code = $1
	`, code).Scan(&score)
	return
}
