package questionnaire

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
)

func questionnaireAnswer(ctx context.Context, req *QuestionnaireStruct) (res QuestionnaireStruct, err error) {
	req.LoanerID = auth.GetUserID(ctx)
	roleID := auth.GetUserRole(ctx)

	if req.LoanerID == 0 {
		return res, ErrSignInRequired
	}
	if roleID != 4 {
		return res, ErrPermissionDeny
	}

	var p prepareArgs
	{
		p.SuggestScore = req.SuggestScore
		p.SuggestGiveScore = req.SuggestGiveScore

		p.income = req.Income
		p.loan = req.Loan
		p.debtPerMonth = req.DebtPerMonth
		p.totalDebt = req.TotalDebt
		p.saving = req.Saving
		p.mortgageSecurities = req.MortgageSecurities

		p.ageCode = req.AgeCode
		p.jobCode = req.JobCode
		p.eduCode = req.EduCode
		p.timeJobCode = req.TimeJobCode
		p.freChangeNameCode = req.FreChangeNameCode
		p.timeOfPhoneNumberCode = req.TimeOfPhoneNumberCode
		p.timeOfNameInHouseParticularCode = req.TimeOfNameInHouseParticularCode
		p.payDebtHistoryCode = req.PayDebtHistoryCode
		p.statusInHouseParticularCode = req.StatusInHouseParticularCode

		p.haveGuarantorCode = req.HaveGuarantorCode
		p.iamGuarantorCode = req.IamGuarantorCode
		p.incomeTrendCode = req.IncomeTrendCode
		p.loanObjectCode = req.LoanObjectCode
		p.provinceCode = req.ProvinceCode
	}

	err = p.prepareData(ctx)
	if err != nil {
		return res, ErrAnswerPrepareData
	}

	req.IncomePerDebt = p.incomePerDebtCode
	req.TotalDebtPerYearIncome = p.totalDebtPerYearIncomeCode
	req.SavingPerLoan = p.savingPerLoanCode
	req.MortgageSecuritiesPerLoan = p.mortgageSecuritiesPerLoanCode

	req.CreditGrade = p.creditGrade
	req.CreditRisk = p.creditRisk
	req.RiskLevel = p.riskLevel
	req.MatrixIndex = p.matrixIndex

	res = *req

	req.StatusID = 1

	/*
		var id int64
		err = dbctx.QueryRow(ctx, `
			insert into questionnaire
				(userID, suggest, suggestScore, suggestGiveScore,
				income, loan, debtPerMonth, totalDebt, saving, mortgageSecurities,
				age, job, edu, timeJob, freChangeName, timeOfPhoneNumber, timeOfNameInHouseParticular, payDebtHistory, statusInHouseParticular,
				incomePerDebt, totalDebtPerYearIncome, savingPerLoan, mortgageSecuritiesPerLoan,
				haveGuarantor, iamGuarantor, incomeTrend, loanObject, provinceCode,
				creditGrade, creditRisk, riskLevel, matrixIndex,
				statusID
				)
			values
				($1,$2, 0, 0,
				$3, $4, $5, $6, $7, $8,
				$9, $10, $11, $12, $13, $14, $15, $16, $17,
				$18, $19, $20, $21,
				$22, $23, $24, $25, $26,
				$27, $28, $29, $30,
				$31
				)
			returning id
			`, req.LoanerID, req.Suggest,
			req.Income, req.Loan, req.DebtPerMonth, req.TotalDebt, req.Saving, req.MortgageSecurities,
			req.AgeCode, req.JobCode, req.EduCode, req.TimeJobCode, req.FreChangeNameCode, req.TimeOfPhoneNumberCode, req.TimeOfNameInHouseParticularCode, req.PayDebtHistoryCode, req.StatusInHouseParticularCode,
			req.IncomePerDebt, req.TotalDebtPerYearIncome, req.SavingPerLoan, req.MortgageSecuritiesPerLoan,
			req.HaveGuarantorCode, req.IamGuarantorCode, req.IncomeTrendCode, req.LoanObjectCode, req.ProvinceCode,
			req.CreditGrade, req.CreditRisk, req.RiskLevel, req.MatrixIndex,
			1).Scan(&id)

	*/

	return
}
