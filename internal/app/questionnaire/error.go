package questionnaire

import (
	"errors"
	"net/http"
)

var (
	ErrSignInRequired         = errors.New("user: not login")
	ErrPermissionDeny         = errors.New("questionnaire: user not permission")
	ErrGetQuestionStatus      = errors.New("can not get questionnaire status")
	ErrThisStatusCanNotUpdate = errors.New("this status loaner can not be update")
	ErrThisStatusCanNotBeSend = errors.New("this status loaner can not be send to verify")
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

// func (s *QuestionnaireStruct) checkNumType()
var (
	ErrIsNilIncomeInput             = errors.New("IncomeInput  is nil")
	ErrIsNilLoanInput               = errors.New("LoanInput is nil")
	ErrIsNilDebtPerMonthInput       = errors.New("DebtPerMonthInput is nil")
	ErrIsNilTotalDebtInput          = errors.New("TotalDebtInput is nil")
	ErrIsNilSavingInput             = errors.New("SavingInput is nil")
	ErrIsNilMortgageSecuritiesInput = errors.New("MortgageSecuritiesInput is nil")

	ErrIncomeMustBeNumber             = errors.New("income must be number")
	ErrLoanMustBeNumber               = errors.New("loan must be number")
	ErrDebtPerMonthMustBeNumber       = errors.New("debt per month must be number")
	ErrTotalDebtMustBeNumber          = errors.New("total debt must be number")
	ErrSavingMustBeNumber             = errors.New("saving must be number")
	ErrMortgageSecuritiesMustBeNumber = errors.New("MortgageSecurities must be number")
)

var (
	ErrGetLoanerAge     = errors.New("func prepare: get loaner age")
	ErrGetProvinceScore = errors.New("func prepare: get province score")
)

// error get score from DB
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
	ErrAnswerPrepareData         = errors.New("questionnaire answer prepare data error")
	ErrQuestionnaireInsert       = errors.New("questionnaire insert questionnaire error")
	ErrQuestionnaireLoanerUpdate = errors.New("questionnaire loaner update data error")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrIncomeMustBeNumber, ErrLoanMustBeNumber, ErrDebtPerMonthMustBeNumber, ErrTotalDebtMustBeNumber, ErrSavingMustBeNumber, ErrMortgageSecuritiesMustBeNumber:
		return http.StatusBadRequest
	case ErrIsNilIncomeInput, ErrIsNilLoanInput, ErrIsNilDebtPerMonthInput, ErrIsNilTotalDebtInput, ErrIsNilSavingInput, ErrIsNilMortgageSecuritiesInput:
		return http.StatusBadRequest
	case ErrSignInRequired, ErrPermissionDeny:
		return http.StatusUnauthorized
	case ErrGetLoanerAge, ErrQuestionnaireLoanerUpdate:
		return http.StatusInternalServerError
	case ErrThisStatusCanNotUpdate, ErrThisStatusCanNotBeSend:
		return http.StatusNotAcceptable

	default:
		return http.StatusInternalServerError
	}
}

func errorToMessage(err error) string {
	switch err {
	case ErrSignInRequired:
		return "ท่านยังไม่ได้เข้าสู่ระบบ"
	case ErrInvalidInputAge:
		return "กรุณาระบุช่วงอายุของท่าน"
	case ErrInvalidInputLoanerID:
		return ""
	case ErrInvalidInputIncome:
		return "กรุณาระบุรายได้ของท่าน และต้องมีค่ามากกว่า 0"
	case ErrInvalidInputLoan:
		return "กรุณาระบุวงเงินที่ท่านต้องการกู้ และต้องมีค่ามากกว่า 0"
	case ErrInvalidInputDebtPerMonth:
		return "กรุณาระบุภาระหนี้สินที่ท่านต้องชำระต่อเดือนในปัจจุบัน และต้องมีค่ามากกว่า 0"
	case ErrInvalidInputTotalDebt:
		return "กรุณาระบุภาระหนี้สินรวมทั้งหมดของท่านในปัจจุบัน และต้องมีค่ามากกว่า 0"
	case ErrInvalidInputSaving:
		return "กรุณาระบุจำนวนเงินออมของท่าน และต้องมีค่ามากกว่า 0"
	case ErrInvalidInputMortgageSecurities:
		return "กรุณาระบุมูลค่าหลักทรัพย์ค้ำประกันของท่าน และต้องมีค่ามากกว่า 0"
	case ErrInvalidInputJob:
		return "กรุณาระบุอาชีพของท่าน"
	case ErrInvalidInputEdu:
		return "กรุณาระบุระดับการศึกษาของท่าน"
	case ErrInvalidInputTimeJob:
		return "กรุณาระบุช่วงระยะเวลาที่ท่านทำงานในอาชีพปัจจุบัน"
	case ErrInvalidInputFreChangeName:
		return "กรุณาระบุจำนวนครั้งที่ท่านได้เปลี่ยนชื่อ"
	case ErrInvalidInputTimeOfPhoneNumber:
		return "กรุณาระบุระยะเวลาการใช้งานของเบอร์โทรศัพท์ของท่าน"
	case ErrInvalidInputTimeOfStayInHouseParticular:
		return "กรุณาระบุระยะเวลาที่ท่านอาศัยอยู่ในทะเบียนบ้านปัจจุบัน"
	case ErrInvalidInputPayDebtHistory:
		return "กรุณาระบุข้อมูลประวัติการชำระหนี้"
	case ErrInvalidInputStatusInHouseParticular:
		return "กรุณาระบุสถานะในทะเบียนบ้านของท่าน"
	case ErrInvalidInputHaveGuarantor:
		return "กรุณาระบุข้อมูลการมีผู้ค้ำประกันของท่าน"
	case ErrInvalidInputIamGuarantor:
		return "กรุณาระบุข้อมูลว่าท่านได้มีการค้ำประกันในผู้อื่นหรือไม่"
	case ErrInvalidInputIncomeTrend:
		return "กรุณาระบุข้อมูลแนวโน้มรายได้ของท่านในอนาคต"
	case ErrInvalidInputLoanObject:
		return "กรุณาระบุข้อมูลวัตถุประสงค์ในการกู้ของท่าน"
	case ErrInvalidInputProvinceCode:
		return "กรุณาเลือกจังหวัดที่ท่านอาศัยอยู่ในปัจจุบัน"

	case ErrIncomeMustBeNumber:
		return "รายได้ต้องเป็นตัวเลขเท่านั้น"

	case ErrLoanMustBeNumber:
		return "วงเงินที่ต้องการกู้ ต้องกรอกข้อมูลเป็นตัวเลขเท่านั้น"
	case ErrDebtPerMonthMustBeNumber:
		return "หนี้สินที่ต้องชำระต่อเดือน ต้องกรอกข้อมูลเป็นตัวเลขเท่านั้น"
	case ErrTotalDebtMustBeNumber:
		return "จำนวนหนี้สินทั้งหมด ต้องกรอกข้อมูลเป็นตัวเลขเท่านั้น"
	case ErrSavingMustBeNumber:
		return "จำนวนเงินออม ต้องกรอกข้อมูลเป็นตัวเลขเท่านั้น"
	case ErrMortgageSecuritiesMustBeNumber:
		return "มูลค่าหลักทรัพย์ค้ำประกัน ต้องกรอกข้อมูลเป็นตัวเลขเท่านั้น"
	case ErrIsNilIncomeInput, ErrIsNilLoanInput, ErrIsNilDebtPerMonthInput, ErrIsNilTotalDebtInput, ErrIsNilSavingInput, ErrIsNilMortgageSecuritiesInput:
		return "is nil"
	case ErrThisStatusCanNotUpdate:
		return "แบบสอบถามของท่านไม่อยู่ในสถานะที่จะแก้ไขได้"
	case ErrThisStatusCanNotBeSend:
		return "แบบสอบถามของท่านไม่อยู่ในสถานะที่จะส่งเพื่อตรวจสอบได้ เนื่องจากท่านอาจได้ส่งข้อมูลเรียบร้อยแล้ว หรือ ท่านยังไม่ได้่ทำแบบสอบถาม"
	default:
		return "internal server error"
	}
}
