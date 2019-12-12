package questionnaire

import (
	"context"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
)

func questionnaireWorkerSend(ctx context.Context, req *QuestionnaireStruct) (res processResponse, err error) {
	/*
		ส่งข้อมูลเพื่อรอการพิจารณาอนุมัติจาก super admin
	*/
	req.WorkerID = auth.GetUserID(ctx)
	roleID := auth.GetUserRole(ctx)

	if req.WorkerID == 0 {
		return res, ErrSignInRequired
	}

	switch roleID {
	case 1, 2, 3:

	default:
		return res, ErrPermissionDeny
	}

	if req.LoanerID == 0 {
		return res, ErrMissingLoanerID
	}

	// พนักงานกำลังตรวจสอบข้อมูล
	req.StatusID = 4

	// check input data type
	err = req.checkNumTypeForWorker()
	if err != nil {
		return res, err
	}

	//prepare data
	var p prepareArgs
	{
		p.SuggestScore = req.SuggestScore
		p.SuggestGiveScore = req.SuggestGiveScore

		p.income = req.IncomeW
		p.loan = req.LoanW
		p.debtPerMonth = req.DebtPerMonthW
		p.totalDebt = req.TotalDebtW
		p.saving = req.SavingW
		p.mortgageSecurities = req.MortgageSecuritiesW

		p.ageCode = req.AgeCodeW
		p.jobCode = req.JobCodeW
		p.eduCode = req.EduCodeW
		p.timeJobCode = req.TimeJobCodeW
		p.freChangeNameCode = req.FreChangeNameCodeW
		p.timeOfPhoneNumberCode = req.TimeOfPhoneNumberCodeW
		p.timeOfNameInHouseParticularCode = req.TimeOfNameInHouseParticularCodeW
		p.payDebtHistoryCode = req.PayDebtHistoryCodeW
		p.statusInHouseParticularCode = req.StatusInHouseParticularCodeW

		p.haveGuarantorCode = req.HaveGuarantorCodeW
		p.iamGuarantorCode = req.IamGuarantorCodeW
		p.incomeTrendCode = req.IncomeTrendCodeW
		p.loanObjectCode = req.LoanObjectCodeW
		p.provinceCode = req.ProvinceCodeW
	}

	err = p.prepareData(ctx)
	if err != nil {
		return res, err
	}

	// set data from prepare to req
	{
		req.IncomePerDebtW = p.incomePerDebtCode
		req.TotalDebtPerYearIncomeW = p.totalDebtPerYearIncomeCode
		req.SavingPerLoanW = p.savingPerLoanCode
		req.MortgageSecuritiesPerLoanW = p.mortgageSecuritiesPerLoanCode

		req.CreditGrade = p.creditGrade
		req.CreditRisk = p.creditRisk
		req.RiskLevel = p.riskLevel
		req.MatrixIndex = p.matrixIndex
	}

	err = req.updateByWorker(ctx)

	if err != nil {
		fmt.Println(err)
		return res, ErrQuestionnaireWorkerVerifyUpdate
	}
	res.Message = "บันทึกการตรวจสอบข้อมูลสำเร็จ"
	return
}
