package questionnaire

import (
	"errors"
	"net/http"
)

var (
	ErrSignInRequired = errors.New("user: not login")
	ErrPermissionDeny = errors.New("questionnaire: user not permission")
)

// error invalid input
var (
	ErrInvalidInputAge                         = errors.New("questionnaire: invalid input age")
	ErrInvalidInputLoanerID                    = errors.New("questionnaire: invalid input loaner id")
	ErrInvalidInputIncome                      = errors.New("questionnaire: invalid input income")
	ErrInvalidInputLoan                        = errors.New("questionnaire: invalid input loan")
	ErrInvalidInputDebtPerMonth                = errors.New("questionnaire: invalid input dept per month")
	ErrInvalidInputTotalDebt                   = errors.New("questionnaire: invalid input total debt")
	ErrInvalidInputSaving                      = errors.New("questionnaire: invalid input saving")
	ErrInvalidInputMortgageSecurities          = errors.New("questionnaire: invalid input MortgageSecurities")
	ErrInvalidInputJob                         = errors.New("questionnaire: job not input")
	ErrInvalidInputEdu                         = errors.New("questionnaire:  not input")
	ErrInvalidInputTimeJob                     = errors.New("questionnaire: TimeJob not input")
	ErrInvalidInputFreChangeName               = errors.New("questionnaire: FreChangeName not input")
	ErrInvalidInputTimeOfPhoneNumber           = errors.New("questionnaire: not input")
	ErrInvalidInputTimeOfStayInHouseParticular = errors.New("questionnaire: TimeOfStayInHouseParticular not input")
	ErrInvalidInputPayDebtHistory              = errors.New("questionnaire: PayDebtHistory not input")
	ErrInvalidInputStatusInHouseParticular     = errors.New("questionnaire: StatusInHouseParticular not input")
	ErrInvalidInputHaveGuarantor               = errors.New("questionnaire: HaveGuarantor not input")
	ErrInvalidInputIamGuarantor                = errors.New("questionnaire: IamGuarantor not input")
	ErrInvalidInputIncomeTrend                 = errors.New("questionnaire: IncomeTrend not input")
	ErrInvalidInputLoanObject                  = errors.New("questionnaire: LoanObject not input")
	ErrInvalidInputProvinceCode                = errors.New("questionnaire: ProvinceCode not input")
)

var (
	ErrGetLoanerAge     = errors.New("func prepare: get loaner age")
	ErrGetProvinceScore = errors.New("func prepare: get province score")
)

var (
	ErrGetAgeScoreFromDB                         = errors.New("error get age score from DB")
	ErrGetJobScoreFromDB                         = errors.New("error get job score from DB")
	ErrGetEduScoreFromDB                         = errors.New("error get edu score from DB")
	ErrGetTimeJobScoreFromDB                     = errors.New("error get time job score from DB")
	ErrGetFreChangeNameScoreFromDB               = errors.New("error get FreChangeName score from DB")
	ErrGetTimeOfNameInHouseParticularScoreFromDB = errors.New("error get TimeOfNameInHouseParticular score from DB")
	ErrGetPayDebtHistoryScoreFromDB              = errors.New("error get PayDebtHistory score from DB")
	ErrGetStatusInHouseParticularScoreFromDB     = errors.New("error get StatusInHouseParticular score from DB")
	ErrGetIncomePerDebtScoreFromDB               = errors.New("error get IncomePerDebt score from DB")
	ErrGetTotalDebtPerYearIncomeScoreFromDB      = errors.New("error get TotalDebtPerYearIncome score from DB")
	ErrGetSavingPerLoanScoreFromDB               = errors.New("error get SavingPerLoan score from DB")
	ErrGetMortgageSecuritiesPerLoanScoreFromDB   = errors.New("error get MortgageSecuritiesPerLoan score from DB")
	ErrGetHaveGuarantorScoreFromDB               = errors.New("error get HaveGuarantor score from DB")
	ErrGetIamGuarantorScoreFromDB                = errors.New("error get IamGuarantor score from DB")
	ErrGetIncomeTrendScoreFromDB                 = errors.New("error get IncomeTrend score from DB")
	ErrGetLoanObjectScoreFromDB                  = errors.New("error get LoanObject score from DB")
	ErrGetProvinceScoreFromDB                    = errors.New("error get Province score from DB")
)

var (
	ErrAnswerPrepareData = errors.New("questionnaire answer prepare data error")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrSignInRequired, ErrPermissionDeny:
		return http.StatusUnauthorized
	case ErrGetLoanerAge:
		return http.StatusInternalServerError

	default:
		return http.StatusInternalServerError
	}
}

func errorToMessage(err error) string {
	switch err {
	case ErrSignInRequired:
		return "ท่านยังไม่ได้เข้าสู่ระบบ"

	default:
		return "internal server error"
	}
}
