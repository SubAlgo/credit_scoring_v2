package questionnaire

import (
	"context"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
	"time"
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

func (res *QuestionnaireStruct) getQuestionnaireData(ctx context.Context, loanerID int64) (err error) {
	var qn questionnaireCheckNull

	err = dbctx.QueryRow(ctx, `
				select 	q.id as questionnaireID, statusID,
						q.loanerID, (select users.name as loanerName from users where users.id = q.loanerID),
						q.updatedBy, (select users.name as updaterName from users where users.id = q.updatedBy),
						q.ApproveBy, (select users.name as updaterName from users where users.id = q.ApproveBy),
						to_char(updatedAt, 'DD Mon YYYY เวลา HH:MM:SS') as updated_at,
						to_char(sendAt, 'DD Mon YYYY เวลา HH:MM:SS') as updated_at,

						creditGrade, creditRisk, riskLevel, matrixIndex,
						approveRate, approveTotal, interest,
						verifyComment, approveComment,
						creditGrade, creditRisk, riskLevel, matrixIndex,
						suggest, suggestScore, suggestGiveScore,

						income, loan, debtPerMonth, totalDebt, saving, mortgageSecurities,
						age, job, edu, timeJob, freChangeName, timeOfPhoneNumber, timeOfNameInHouseParticular, payDebtHistory, statusInHouseParticular,
						incomePerDebt, totalDebtPerYearIncome, savingPerLoan, mortgageSecuritiesPerLoan,
						haveGuarantor, iamGuarantor, incomeTrend, loanObject, provinceCode,

						incomeW, loanW, debtPerMonthW, totalDebtW, savingW, mortgageSecurities,
						ageW, jobW, eduW, timeJobW, freChangeNameW, timeOfPhoneNumberW, timeOfNameInHouseParticularW, payDebtHistoryW, statusInHouseParticular,
						incomePerDebtW, totalDebtPerYearIncomeW, savingPerLoanW, mortgageSecuritiesPerLoanW,
						haveGuarantorW, iamGuarantorW, incomeTrendW, loanObjectW, provinceCodeW

				from questionnaire as q
				left join users as u on u.id = q.loanerID 
				where loanerID = $1;
			`, loanerID).Scan(&res.ID, &res.StatusID,
		&res.LoanerID, &res.LoanerName,
		&res.UpdatedBy, &res.UpdatedByName,
		&qn.ApproveBy, &qn.ApproveName,
		&res.UpdateAtStr,
		&res.SendAtStr,

		&res.CreditGrade, &res.CreditRisk, &res.RiskLevel, &res.MatrixIndex,
		&qn.ApproveRate, &qn.ApproveTotal, &qn.Interest,
		&qn.VerifyComment, &qn.ApproveComment,
		&res.CreditGrade, &res.CreditRisk, &res.RiskLevel, &res.MatrixIndex,
		&res.Suggest, &res.SuggestScore, &res.SuggestGiveScore,

		&res.IncomeInput, &res.LoanInput, &res.DebtPerMonthInput, &res.TotalDebtInput, &res.SavingInput, &res.MortgageSecuritiesInput,
		&res.AgeCode, &res.JobCode, &res.EduCode, &res.TimeJobCode, &res.FreChangeNameCode, &res.TimeOfPhoneNumberCode, &res.TimeOfNameInHouseParticularCode, &res.PayDebtHistoryCode, &res.StatusInHouseParticularCode,
		&res.IncomePerDebt, &res.TotalDebtPerYearIncome, &res.SavingPerLoan, &res.MortgageSecuritiesPerLoan,
		&res.HaveGuarantorCode, &res.IamGuarantorCode, &res.IncomeTrendCode, &res.LoanObjectCode, &res.ProvinceCode,

		&res.IncomeInputW, &res.LoanInputW, &res.DebtPerMonthInputW, &res.TotalDebtInputW, &res.SavingInputW, &res.MortgageSecuritiesInputW,
		&res.AgeCodeW, &res.JobCodeW, &res.EduCodeW, &res.TimeJobCodeW, &res.FreChangeNameCodeW, &res.TimeOfPhoneNumberCodeW, &res.TimeOfNameInHouseParticularCodeW, &res.PayDebtHistoryCodeW, &res.StatusInHouseParticularCodeW,
		&res.IncomePerDebtW, &res.TotalDebtPerYearIncomeW, &res.SavingPerLoanW, &res.MortgageSecuritiesPerLoanW,
		&res.HaveGuarantorCodeW, &res.IamGuarantorCodeW, &res.IncomeTrendCodeW, &res.LoanObjectCodeW, &res.ProvinceCodeW)

	if err != nil {
		fmt.Println(err)
		return ErrQuestionnaireSelectData
	}

	res.ApproveBy = qn.ApproveBy.Int64
	res.ApproveName = qn.ApproveName.String
	res.ApproveRate = qn.ApproveRate.Float64
	res.ApproveTotal = qn.ApproveTotal.Float64
	res.Interest = qn.Interest.Float64
	res.VerifyComment = qn.VerifyComment.String
	res.ApproveComment = qn.ApproveComment.String

	return
}

func (req *QuestionnaireStruct) updateByLoaner(ctx context.Context) (err error) {
	updatedAt := time.Now()
	_, err = dbctx.Exec(ctx, `
		update questionnaire 
		set updatedBy = $1,
			suggest = $2, 
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
			statusID = $31,

			updatedAt = $32,
			sendAt = $32

		where loanerID = $1
	`, req.LoanerID, req.Suggest, req.Income, req.Loan, req.DebtPerMonth, req.TotalDebt, req.Saving, req.MortgageSecurities,
		req.AgeCode, req.JobCode, req.EduCode, req.TimeJobCode, req.FreChangeNameCode, req.TimeOfPhoneNumberCode, req.TimeOfNameInHouseParticularCode, req.PayDebtHistoryCode, req.StatusInHouseParticularCode,
		req.IncomePerDebt, req.TotalDebtPerYearIncome, req.SavingPerLoan, req.MortgageSecuritiesPerLoan,
		req.HaveGuarantorCode, req.IamGuarantorCode, req.IncomeTrendCode, req.LoanObjectCode, req.ProvinceCode,
		req.CreditGrade, req.CreditRisk, req.RiskLevel, req.MatrixIndex, req.StatusID, updatedAt)
	return
}

func (req *QuestionnaireStruct) updateByWorker(ctx context.Context) (err error) {
	_, err = dbctx.Exec(ctx, `
		update questionnaire
		set updatedBy = $2,

			suggestScore = $3,
			suggestGiveScore = $4,
			
			incomeW = $5, 
			loanW = $6, 
			debtPerMonthW = $7, 
			totalDebtW = $8, 
			savingW = $9, 
			mortgageSecuritiesW = $10,

			ageW = $11, 
			jobW = $12, 
			eduW = $13,
			timeJobW = $14,
			freChangeNameW = $15,
			timeOfPhoneNumberW = $16,
			timeOfNameInHouseParticularW = $17, 
			payDebtHistoryW = $18,
			statusInHouseParticularW = $19,
				
			incomePerDebtW = $20, 
			totalDebtPerYearIncomeW = $21, 
			savingPerLoanW = $22, 
			mortgageSecuritiesPerLoanW = $23,
				
			haveGuarantorW =$24, 
			iamGuarantorW =$25, 
			incomeTrendW =$26, 
			loanObjectW = $27,
			provinceCodeW = $28,

			creditGrade = $29,
			creditRisk =  $30,
			riskLevel =  $31,
			matrixIndex =  $32,

			verifyComment = $33,
			statusID = $34

		where loanerID = $1
	`, req.LoanerID, req.WorkerID,                                                                                                                                                                                       //1-2
		req.SuggestScore, req.SuggestGiveScore,                                                                                                                                                                          //3-4
		req.IncomeW, req.LoanW, req.DebtPerMonthW, req.TotalDebtW, req.SavingW, req.MortgageSecuritiesW,                                                                                                                 //5-10
		req.AgeCodeW, req.JobCodeW, req.EduCodeW, req.TimeJobCodeW, req.FreChangeNameCodeW, req.TimeOfPhoneNumberCodeW, req.TimeOfNameInHouseParticularCodeW, req.PayDebtHistoryCodeW, req.StatusInHouseParticularCodeW, //11-19
		req.IncomePerDebtW, req.TotalDebtPerYearIncomeW, req.SavingPerLoanW, req.MortgageSecuritiesPerLoanW,                                                                                                             //20-23
		req.HaveGuarantorCodeW, req.IamGuarantorCodeW, req.IncomeTrendCodeW, req.LoanObjectCodeW, req.ProvinceCodeW,                                                                                                     //24-28
		req.CreditGrade, req.CreditRisk, req.RiskLevel, req.MatrixIndex,                                                                                                                                                 //
		req.VerifyComment, req.StatusID)
	return
}
