package questionnaire

import (
	"context"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
	"time"
)

func questionnaireAnswer(ctx context.Context, req *QuestionnaireStruct) (res processResponse, err error) {
	req.LoanerID = auth.GetUserID(ctx)
	roleID := auth.GetUserRole(ctx)

	if req.LoanerID == 0 {
		return res, ErrSignInRequired
	}
	if roleID != 4 {
		return res, ErrPermissionDeny
	}

	// set status id
	req.StatusID = 1

	// check input data type
	err = req.checkNumType()
	if err != nil {
		return res, err
	}

	//prepare data
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
		return res, err
	}

	// set data from prepare to req
	{
		req.IncomePerDebt = p.incomePerDebtCode
		req.TotalDebtPerYearIncome = p.totalDebtPerYearIncomeCode
		req.SavingPerLoan = p.savingPerLoanCode
		req.MortgageSecuritiesPerLoan = p.mortgageSecuritiesPerLoanCode

		req.CreditGrade = p.creditGrade
		req.CreditRisk = p.creditRisk
		req.RiskLevel = p.riskLevel
		req.MatrixIndex = p.matrixIndex
	}

	var id int64

	updateAt := time.Now()
	var defaultFloatValue float64 = 0
	defaultStringValue := "-"
	err = dbctx.QueryRow(ctx, `
				insert into questionnaire
				(loanerID, updatedBy,
				suggest, suggestW, suggestScore, suggestGiveScore,
				income, loan, debtPerMonth, totalDebt, saving, mortgageSecurities,
				age, job, edu, timeJob, freChangeName, timeOfPhoneNumber, timeOfNameInHouseParticular, payDebtHistory, statusInHouseParticular,
				incomePerDebt, totalDebtPerYearIncome, savingPerLoan, mortgageSecuritiesPerLoan,
				haveGuarantor, iamGuarantor, incomeTrend, loanObject, provinceCode,

				incomeW, loanW, debtPerMonthW, totalDebtW, savingW, mortgageSecuritiesW,
				ageW, jobW, eduW, timeJobW, freChangeNameW, timeOfPhoneNumberW, timeOfNameInHouseParticularW, payDebtHistoryW, statusInHouseParticularW,
				incomePerDebtW, totalDebtPerYearIncomeW, savingPerLoanW, mortgageSecuritiesPerLoanW,
				haveGuarantorW, iamGuarantorW, incomeTrendW, loanObjectW, provinceCodeW,

				creditGrade, creditRisk, riskLevel, matrixIndex,
				statusID, updatedAt, sendAt,
				approveRate, approveTotal, interest,
				verifyComment, approveComment
				)
				values
				($1, $1,
				$2, $2, $3, $4,
				$5, $6, $7, $8,$9, $10,
				$11, $12, $13, $14, $15, $16, $17,$18, $19,
				$20, $21,$22, $23,
				$24, $25, $26,$27, $28,

				$5, $6, $7, $8,$9, $10,
				$11, $12, $13, $14, $15, $16, $17,$18, $19,
				$20, $21,$22, $23,
				$24, $25, $26,$27, $28,

				$29, $30, $31, $32,
				$33, $34, $34,
				$35, $35, $35,
				$36, $36)
				returning id
				`, req.LoanerID, req.Suggest, req.SuggestScore, req.SuggestGiveScore,
		req.Income, req.Loan, req.DebtPerMonth, req.TotalDebt, req.Saving, req.MortgageSecurities,
		req.AgeCode, req.JobCode, req.EduCode, req.TimeJobCode, req.FreChangeNameCode, req.TimeOfPhoneNumberCode, req.TimeOfNameInHouseParticularCode, req.PayDebtHistoryCode, req.StatusInHouseParticularCode,
		req.IncomePerDebt, req.TotalDebtPerYearIncome, req.SavingPerLoan, req.MortgageSecuritiesPerLoan,
		req.HaveGuarantorCode, req.IamGuarantorCode, req.IncomeTrendCode, req.LoanObjectCode, req.ProvinceCode,
		req.CreditGrade, req.CreditRisk, req.RiskLevel, req.MatrixIndex,
		req.StatusID, updateAt,
		defaultFloatValue,
		defaultStringValue).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return res, ErrQuestionnaireInsert
	}

	res.Message = "'แบบประเมินขออนุมัติสินเชื่อของท่านได้ถูกบันทึกเรียบร้อยแล้ว กรุณาตรวจสอบความถูกต้องของข้อมูล เพื่อส่งแบบประเมินขออนุมัติสินเชื่อแก่เจ้าหน้าที่ ในหน้าถัดไป'"

	return
}
