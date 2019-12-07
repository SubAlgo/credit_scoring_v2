package questionnaire

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
)

func questionnaireLoanerUpdate(ctx context.Context, req *QuestionnaireStruct) (res processResponse, err error) {
	req.LoanerID = auth.GetUserID(ctx)
	roleID := auth.GetUserRole(ctx)

	if req.LoanerID == 0 {
		return res, ErrSignInRequired
	}
	if roleID != 4 {
		return res, ErrPermissionDeny
	}

	// check questionnaire status
	req.StatusID, err = checkQuestionnaireStatus(ctx, req.LoanerID)
	if err != nil {
		return res, ErrGetQuestionStatus
	}

	//check questionnaire status
	if req.StatusID != 1 {
		return res, ErrThisStatusCanNotUpdate
	}

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

	// set questionnaire status id
	req.StatusID = 1

	err = req.updateByLoaner(ctx)

	if err != nil {
		return res, ErrQuestionnaireLoanerUpdate
	}

	res.Message = "แก้ไขข้อมูลสำเร็จ"
	return
}
