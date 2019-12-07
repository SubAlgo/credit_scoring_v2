package questionnaire

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

// questionnaire insert
func (q *QuestionnaireStruct) insert() () {

}

func checkQuestionnaireStatus(ctx context.Context, loanerID int64) (questionnaireStatus int, err error) {
	err = dbctx.QueryRow(ctx, `
			select statusID
			from questionnaire
			where loanerID = $1
	`, loanerID).Scan(&questionnaireStatus)
	return
}

func (req *QuestionnaireStruct) updateByLoaner(ctx context.Context) (err error) {
	_, err = dbctx.Exec(ctx, `
		update questionnaire 
		set suggest = $2, 
			income = $3, 
			loan = $4, 
			debtPerMonth = $5, 
			totalDebt = $6, 
			saving = $7, 
			mortgageSecurities = $8,
		
			age = $9, 
			job = $10, 
			edu = $11, 
			timeJob = $12, 
			freChangeName = $13, 
			timeOfPhoneNumber = $14, 
			timeOfNameInHouseParticular = $15, 
			payDebtHistory = $16, 
			statusInHouseParticular = $17,
				
			incomePerDebt = $18, 
			totalDebtPerYearIncome = $19, 
			savingPerLoan = $20, 
			mortgageSecuritiesPerLoan = $21,
				
			haveGuarantor = $22, 
			iamGuarantor = $23, 
			incomeTrend = $24, 
			loanObject = $25, 
			provinceCode = $26,

			incomeW = $3, 
			loanW = $4, 
			debtPerMonthW = $5, 
			totalDebtW = $6, 
			savingW = $7, 
			mortgageSecuritiesW = $8,
		
			ageW = $9, 
			jobW = $10, 
			eduW = $11, 
			timeJobW = $12, 
			freChangeNameW = $13, 
			timeOfPhoneNumberW = $14, 
			timeOfNameInHouseParticularW = $15, 
			payDebtHistoryW = $16, 
			statusInHouseParticularW = $17,
				
			incomePerDebtW = $18, 
			totalDebtPerYearIncomeW = $19, 
			savingPerLoanW = $20, 
			mortgageSecuritiesPerLoanW = $21,
				
			haveGuarantorW = $22, 
			iamGuarantorW = $23, 
			incomeTrendW = $24, 
			loanObjectW = $25, 
			provinceCodeW = $26,

			creditGrade = $27 ,
			creditRisk =  $28,
			riskLevel =  $29,
			matrixIndex =  $30,
			statusID = $31

		where loanerID = $1
	`, req.LoanerID, req.Suggest, req.Income, req.Loan, req.DebtPerMonth, req.TotalDebt, req.Saving, req.MortgageSecurities,
		req.AgeCode, req.JobCode, req.EduCode, req.TimeJobCode, req.FreChangeNameCode, req.TimeOfPhoneNumberCode, req.TimeOfNameInHouseParticularCode, req.PayDebtHistoryCode, req.StatusInHouseParticularCode,
		req.IncomePerDebt, req.TotalDebtPerYearIncome, req.SavingPerLoan, req.MortgageSecuritiesPerLoan,
		req.HaveGuarantorCode, req.IamGuarantorCode, req.IncomeTrendCode, req.LoanObjectCode, req.ProvinceCode,
		req.CreditGrade, req.CreditRisk, req.RiskLevel, req.MatrixIndex, req.StatusID)
	return
}
