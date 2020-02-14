package questionnaire

import (
	"time"
)

type QuestionnaireStruct struct {
	ID            int64  `json:"id"`
	LoanerID      int64  `json:"loanerID"`
	LoanerName    string `json:"loanerName"`
	WorkerID      int64
	UpdatedBy     int64     `json:"updatedBy"`
	UpdatedByName string    `json:"updatedByName"`
	ApproveBy     int64     `json:"approveBy"`
	ApproveName   string    `json:"approveName"`
	SendAt        time.Time `json:"sendAt"`
	SendAtStr     string    `json:"send_at_str"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UpdateAtStr   string    `json:"update_at_str"`

	StatusID int `json:"statusID"`

	ApproveRate  float64 `json:"approveRate"`
	ApproveTotal float64 `json:"approveTotal"`
	Interest     float64 `json:"interest"`

	VerifyComment  string `json:"verifyComment"`
	ApproveComment string `json:"approveComment"`

	CreditGrade string `json:"creditGrade"`
	CreditRisk  string `json:"creditRisk"`
	RiskLevel   string `json:"riskLevel"`
	MatrixIndex string `json:"matrixIndex"`

	// -- ข้อมูลผู้แนะนำ
	Suggest          string `json:"suggest"`
	SuggestW         string `json:"suggestW"`
	SuggestScore     int    `json:"suggestScore"`
	SuggestGiveScore int    `json:"suggestGiveScore"`

	//-- ข้อมูลตัวเลขกรอกมือ
	IncomeInput             interface{} `json:"incomeInput"`
	LoanInput               interface{} `json:"loanInput"`
	DebtPerMonthInput       interface{} `json:"debtPerMonthInput"`
	TotalDebtInput          interface{} `json:"totalDebtInput"`
	SavingInput             interface{} `json:"savingInput"`
	MortgageSecuritiesInput interface{} `json:"mortgageSecuritiesInput"`

	//-- ข้อมูลตัวเลขกรอกมือ
	IncomeInputW             interface{} `json:"incomeInputW"`
	LoanInputW               interface{} `json:"loanInputW"`
	DebtPerMonthInputW       interface{} `json:"debtPerMonthInputW"`
	TotalDebtInputW          interface{} `json:"totalDebtInputW"`
	SavingInputW             interface{} `json:"savingInputW"`
	MortgageSecuritiesInputW interface{} `json:"mortgageSecuritiesInputW"`

	//-- ข้อมูลตัวเลขกรอกมือ
	Income             float64 `json:"income"`
	Loan               float64 `json:"loan"`
	DebtPerMonth       float64 `json:"debtPerMonth"`
	TotalDebt          float64 `json:"totalDebt"`
	Saving             float64 `json:"saving"`
	MortgageSecurities float64 `json:"mortgageSecurities"`

	//-- ข้อมูลตัวเลือก (option)
	AgeCode                         string `json:"ageCode"` //code
	JobCode                         string `json:"jobCode"`
	EduCode                         string `json:"eduCode"`
	TimeJobCode                     string `json:"timeJobCode"`
	FreChangeNameCode               string `json:"freChangeNameCode"`
	TimeOfPhoneNumberCode           string `json:"timeOfPhoneNumberCode"`
	TimeOfNameInHouseParticularCode string `json:"timeOfNameInHouseParticularCode"`
	PayDebtHistoryCode              string `json:"payDebtHistoryCode"`
	StatusInHouseParticularCode     string `json:"statusInHouseParticularCode"`

	IncomePerDebt             string `json:"incomePerDebtCode"`
	TotalDebtPerYearIncome    string `json:"totalDebtPerYearIncomeCode"`
	SavingPerLoan             string `json:"savingPerLoanCode"`
	MortgageSecuritiesPerLoan string `json:"mortgageSecuritiesPerLoanCode"`

	HaveGuarantorCode string `json:"haveGuarantorCode"`
	IamGuarantorCode  string `json:"iamGuarantorCode"`
	IncomeTrendCode   string `json:"incomeTrendCode"`
	LoanObjectCode    string `json:"loanObjectCode"`
	ProvinceCode      string `json:"provinceCode"`

	//-- ส่วนผู้ตรวจสอบแก้ไข
	IncomeW             float64 `json:"incomeW"`
	LoanW               float64 `json:"loanW"`
	DebtPerMonthW       float64 `json:"debtPerMonthW"`
	TotalDebtW          float64 `json:"totalDebtW"`
	SavingW             float64 `json:"	savingW"`
	MortgageSecuritiesW float64 `json:"mortgageSecuritiesW"`

	//-- ข้อมูลตัวเลือก (option)
	AgeCodeW                         string `json:"ageCodeW"`
	JobCodeW                         string `json:"jobCodeW"`
	EduCodeW                         string `json:"eduCodeW"`
	TimeJobCodeW                     string `json:"timeJobCodeW"`
	FreChangeNameCodeW               string `json:"freChangeNameCodeW"`
	TimeOfPhoneNumberCodeW           string `json:"timeOfPhoneNumberCodeW"`
	TimeOfNameInHouseParticularCodeW string `json:"timeOfNameInHouseParticularCodeW"`
	PayDebtHistoryCodeW              string `json:"payDebtHistoryCodeW"`
	StatusInHouseParticularCodeW     string `json:"statusInHouseParticularCodeW"`

	IncomePerDebtW             string
	TotalDebtPerYearIncomeW    string
	SavingPerLoanW             string
	MortgageSecuritiesPerLoanW string

	HaveGuarantorCodeW string `json:"haveGuarantorCodeW"`
	IamGuarantorCodeW  string `json:"iamGuarantorCodeW"`
	IncomeTrendCodeW   string `json:"incomeTrendCodeW"`
	LoanObjectCodeW    string `json:"loanObjectCodeW"`
	ProvinceCodeW      string `json:"provinceCodeW"`
}

type questionnaireCheckNull struct {
	ID          NullInt
	LoanerID    NullInt
	UpdatedBy   NullInt
	ApproveBy   NullInt
	ApproveName NullString

	StatusID  NullInt
	SendAt    NullTime
	UpdatedAt NullTime

	ApproveRate  NullFloat64
	ApproveTotal NullFloat64
	Interest     NullFloat64

	VerifyComment  NullString
	ApproveComment NullString

	CreditGrade NullString
	CreditRisk  NullString
	RiskLevel   NullString
	MatrixIndex NullString

	// -- ส่วนผู้กู้กรอก
	Suggest          NullString
	SuggestScore     NullInt
	SuggestGiveScore NullInt

	//-- ข้อมูลตัวเลขกรอกมือ
	Income             NullFloat64
	Loan               NullFloat64
	DebtPerMonth       NullFloat64
	TotalDebt          NullFloat64
	Saving             NullFloat64
	MortgageSecurities NullFloat64

	//-- ข้อมูลตัวเลือก (option)
	Age                         NullString
	Job                         NullString
	Edu                         NullString
	TimeJob                     NullString
	FreChangeName               NullString
	TimeOfPhoneNumber           NullString
	TimeOfNameInHouseParticular NullString
	PayDebtHistory              NullString
	StatusInHouseParticular     NullString

	IncomePerDebt             NullString
	TotalDebtPerYearIncome    NullString
	SavingPerLoan             NullString
	MortgageSecuritiesPerLoan NullString

	HaveGuarantor NullString
	IamGuarantor  NullString
	IncomeTrend   NullString
	LoanObject    NullString
	ProvinceCode  NullString

	//-- ส่วนผู้ตรวจสอบแก้ไข
	IncomeW             NullFloat64
	LoanW               NullFloat64
	DebtPerMonthW       NullFloat64
	TotalDebtW          NullFloat64
	SavingW             NullFloat64
	MortgageSecuritiesW NullFloat64

	//-- ข้อมูลตัวเลือก (option)
	AgeW                         NullString
	JobW                         NullString
	EduW                         NullString
	TimeJobW                     NullString
	FreChangeNameW               NullString
	TimeOfPhoneNumberW           NullString
	TimeOfNameInHouseParticularW NullString
	PayDebtHistoryW              NullString
	StatusInHouseParticularW     NullString

	IncomePerDebtW             NullString
	TotalDebtPerYearIncomeW    NullString
	SavingPerLoanW             NullString
	MortgageSecuritiesPerLoanW NullString

	HaveGuarantorW NullString
	IamGuarantorW  NullString
	IncomeTrendW   NullString
	LoanObjectW    NullString
	ProvinceCodeW  NullString
}

type processResponse struct {
	Message string `json:"message"`
}
