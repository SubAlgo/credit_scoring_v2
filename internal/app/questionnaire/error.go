package questionnaire

import (
	"errors"
	"net/http"
)

var (
	ErrSignInRequired                         = errors.New("user: not login")
	ErrPermissionDeny                         = errors.New("questionnaire: user not permission")
	ErrGetQuestionStatus                      = errors.New("can not get questionnaire status")
	ErrThisStatusCanNotUpdate                 = errors.New("this status loaner can not be update")
	ErrThisStatusCanNotBeSend                 = errors.New("this status loaner can not be send to verify")
	ErrMissingLoanerID                        = errors.New("missing loaner id")
	ErrQuestionnaireSelectDataMissingLoanerID = errors.New("questionnaire select data missing loaner id")
)

var (
	ErrQuestionnaireGetApproveResultDB = errors.New("get approve result: DB")
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
	ErrIsNilIncomeInput             = errors.New("IncomeInput is nil")
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

// get questionnaire get status
var (
	ErrQuestionnaireGetStatus = errors.New("questionnaireCheckStatus: get status error")
)

// questionnaireCountList
var (
	ErrGetTotalNewLoaner         = errors.New("questionnaireCountList: get total new loaner")
	ErrGetTotalVerifyLoaner      = errors.New("questionnaireCountList: get total on verify loaner")
	ErrGetTotalWaitApproveLoaner = errors.New("questionnaireCountList: get total wait approve loaner")
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
	ErrAnswerPrepareData                 = errors.New("questionnaire answer prepare data error")
	ErrQuestionnaireInsert               = errors.New("questionnaire insert questionnaire error")
	ErrQuestionnaireLoanerUpdate         = errors.New("questionnaire loaner update data error")
	ErrQuestionnaireWorkerVerifyUpdate   = errors.New("questionnaire worker verify update db error")
	ErrQuestionnaireSelectData           = errors.New("questionnaire select data ")
	ErrQuestionnaireStatusIDNotAvailable = errors.New("questionnaire required status id")
	ErrApproveRateNotAvailable           = errors.New("questionnaire approve rate not available")
	ErrQuestionnaireInterestNotAvailable = errors.New("questionnaire interest rate not available")
	ErrGetLoan                           = errors.New("questionnaire get loan data")
	ErrQuestionnaireApprove              = errors.New("questionnaire approve update db")
	ErrYourAgeCanNotLowerThan20          = errors.New("questionnaire: age can not lower than 20")
)

// get list
var (
	ErrQuestionnaireGetListNewLoaner   = errors.New("questionnaire get list new loaner")
	ErrSendBackToVerifyUpdateDB        = errors.New("questionnaireSendBackToVerify: update database error")
	ErrSendBackToVerifyLoanerIDRequest = errors.New("questionnaireSendBackToVerify: request loaner id")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrIncomeMustBeNumber, ErrLoanMustBeNumber, ErrDebtPerMonthMustBeNumber, ErrTotalDebtMustBeNumber, ErrSavingMustBeNumber, ErrMortgageSecuritiesMustBeNumber, ErrQuestionnaireSelectDataMissingLoanerID:
		return http.StatusBadRequest
	case ErrIsNilIncomeInput, ErrSendBackToVerifyLoanerIDRequest, ErrIsNilLoanInput, ErrIsNilDebtPerMonthInput, ErrIsNilTotalDebtInput, ErrIsNilSavingInput, ErrIsNilMortgageSecuritiesInput, ErrQuestionnaireStatusIDNotAvailable, ErrApproveRateNotAvailable, ErrQuestionnaireInterestNotAvailable:
		return http.StatusBadRequest
	case ErrSignInRequired, ErrPermissionDeny:
		return http.StatusUnauthorized
	case ErrGetLoanerAge, ErrSendBackToVerifyUpdateDB, ErrQuestionnaireLoanerUpdate, ErrQuestionnaireWorkerVerifyUpdate, ErrQuestionnaireSelectData, ErrQuestionnaireGetStatus:
		return http.StatusInternalServerError
	case ErrYourAgeCanNotLowerThan20:
		return http.StatusBadRequest
	case ErrGetTotalNewLoaner, ErrGetTotalVerifyLoaner, ErrGetTotalWaitApproveLoaner:
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
	case ErrPermissionDeny:
		return "ท่านไม่มีสิทธิเข้าใช้งานการใช้งานนี้ได้"
	case ErrInvalidInputAge:
		return "กรุณาระบุช่วงอายุของท่าน"
	case ErrMissingLoanerID:
		return "กรุณาระบุระหัสของผู้ขอสินเชื่อ"
	case ErrInvalidInputLoanerID:
		return "need loaner id"
	case ErrInvalidInputIncome:
		return "กรุณาระบุรายได้ และข้อมูลต้องมีค่ามากกว่า 0"
	case ErrInvalidInputLoan:
		return "กรุณาระบุจำนวนสินเชื่อที่ต้องการ และข้อมูลต้องมีค่ามากกว่า 0"
	case ErrInvalidInputDebtPerMonth:
		return "กรุณาระบุภาระหนี้สินที่ต้องชำระต่อเดือนในปัจจุบัน โดยข้อมูลไม่สามารถมีค่าน้อยกว่า 0"
	case ErrInvalidInputTotalDebt:
		return "กรุณาระบุภาระหนี้สินรวมทั้งหมดของในปัจจุบัน โดยข้อมูลไม่สามารถมีค่าน้อยกว่า 0"
	case ErrInvalidInputSaving:
		return "กรุณาระบุจำนวนเงินออม โดยข้อมูลไม่สามารถมีค่าน้อยกว่า 0"
	case ErrInvalidInputMortgageSecurities:
		return "กรุณาระบุมูลค่าหลักทรัพย์ค้ำประกัน โดยข้อมูลไม่สามารถมีค่าน้อยกว่า 0"
	case ErrInvalidInputJob:
		return "กรุณาระบุข้อมูลอาชีพ"
	case ErrInvalidInputEdu:
		return "กรุณาระบุข้อมูลระดับการศึกษา"
	case ErrInvalidInputTimeJob:
		return "กรุณาระบุข้อมูลระยะประสบการณ์ของอาชีพที่ทำในปัจจุบัน"
	case ErrInvalidInputFreChangeName:
		return "กรุณาระบุข้อมูลจำนวนครั้งในการเปลี่ยนชื่อ"
	case ErrInvalidInputTimeOfPhoneNumber:
		return "กรุณาระบุข้อมูลช่วงจำนวนปีที่ใช้งานเบอร์โทรศัพท์หมายเลขปัจจุบัน"
	case ErrInvalidInputTimeOfStayInHouseParticular:
		return "กรุณาระบุข้อมูลระยะเวลาที่อาศัยอยู่ในทะเบียนบ้านปัจจุบัน"
	case ErrInvalidInputPayDebtHistory:
		return "กรุณาระบุข้อมูลประวัติการชำระหนี้"
	case ErrInvalidInputStatusInHouseParticular:
		return "กรุณาระบุข้อมูลสถานะในทะเบียนบ้าน"
	case ErrInvalidInputHaveGuarantor:
		return "กรุณาระบุข้อมูลการมีผู้ค้ำประกัน"
	case ErrInvalidInputIamGuarantor:
		return "กรุณาระบุข้อมูลสถานะการค้ำประกันให้บุคคลอื่น"
	case ErrInvalidInputIncomeTrend:
		return "กรุณาระบุข้อมูลความคิดเห็นต่อแนวโน้มรายได้ในอนาคตของตัวผู้ขอสินเชื่อ"
	case ErrInvalidInputLoanObject:
		return "กรุณาระบุข้อมูลวัตถุประสงค์ในการขอสินเชื่อ"
	case ErrInvalidInputProvinceCode:
		return "กรุณาเลือกจังหวัดที่ผู้ขอสินเชื่ออาศัยอยู่ในปัจจุบัน"

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
	case ErrIsNilIncomeInput:
		return "IncomeInput is nil value"
	case ErrIsNilLoanInput:
		return "LoanInput is nil value"
	case ErrIsNilDebtPerMonthInput:
		return "DebtPerMonthInput is nil value"
	case ErrIsNilTotalDebtInput:
		return "TotalDebtInput is nil value"
	case ErrIsNilSavingInput:
		return "SavingInput is nil value"
	case ErrIsNilMortgageSecuritiesInput:
		return "MortgageSecuritiesInput is nil value"
	case ErrThisStatusCanNotUpdate:
		return "แบบสอบถามของท่านไม่อยู่ในสถานะที่จะแก้ไขได้"
	case ErrThisStatusCanNotBeSend:
		return "แบบสอบถามของท่านไม่อยู่ในสถานะที่จะส่งเพื่อตรวจสอบได้ เนื่องจากท่านอาจได้ส่งข้อมูลเรียบร้อยแล้ว หรือ ท่านยังไม่ได้่ทำแบบสอบถาม"
	case ErrQuestionnaireWorkerVerifyUpdate:
		return "internal server error (verify update db)"
	case ErrQuestionnaireSelectData:
		return "เกิดข้อผิดพลาดในการดึงข้อมูล"
	case ErrQuestionnaireSelectDataMissingLoanerID:
		return "questionnaire select data -> กรุณาระบุรหัสผู้กู้"
	case ErrQuestionnaireStatusIDNotAvailable:
		return "status id not available"
	case ErrApproveRateNotAvailable:
		return "อัตราการอนุมัติสินเชื่อต้องมีค่าระหว่าง 1 - 100 เท่านั้น"
	case ErrQuestionnaireInterestNotAvailable:
		return "อัตราดอกเบี้ยต้องมีค่าไม่น้อยกว่า 0"
	case ErrQuestionnaireApprove:
		return "internal server error (approve)"
	case ErrQuestionnaireGetListNewLoaner:
		return "internal server error (questionnaireGetListNewLoaner)"
	case ErrQuestionnaireGetStatus:
		return "internal server error (get questionnaire status)"
	case ErrQuestionnaireGetApproveResultDB:
		return "internal server error (get approve result db)"
	case ErrSendBackToVerifyLoanerIDRequest:
		return "ท่านไม่ได้ระบุรหัสผู้กู้"
	case ErrYourAgeCanNotLowerThan20:
		return "ไม่สามารถอนุมัติสินเชื่อให้ผู้มีอายุต่ำกว่า 20ปีได้"
	case ErrGetTotalNewLoaner:
		return "เกิดข้อผิดพลาดในการนับจำนวนผู้กู้ (รายชื่อใหม่)"
	case ErrGetTotalVerifyLoaner:
		return "เกิดข้อผิดพลาดในการนับจำนวนผู้กู้ (รายชื่อที่อยู่ระหว่างการตรวจสอบ)"
	case ErrGetTotalWaitApproveLoaner:
		return "เกิดข้อผิดพลาดในการนับจำนวนผู้กู้ (รายชื่อที่รอการพิจารณาอนุมัติ)"
	default:
		return "internal server error"
	}
}
