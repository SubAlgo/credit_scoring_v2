package questionnaire

import (
	"database/sql"
	"reflect"
	"time"
)

type QuestionnaireStruct struct {
	ID        int64     `json:"id"`
	LoanerID  int64     `json:"loanerID"`
	UpdatedBy int64     `json:"updatedBy"`
	ApproveBy int64     `json:"approveBy"`
	SendAt    time.Time `json:"sendAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	StatusID  int       `json:"statusID"`

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
	AgeCodeW                         string `json:"ageW"`
	JobCodeW                         string `json:"JobW"`
	EduCodeW                         string `json:"eduW"`
	TimeJobCodeW                     string `json:"timeJobW"`
	FreChangeNameCodeW               string `json:"freChangeNameW"`
	TimeOfPhoneNumberCodeW           string `json:"timeOfPhoneNumberW"`
	TimeOfNameInHouseParticularCodeW string `json:"timeOfNameInHouseParticularW"`
	PayDebtHistoryCodeW              string `json:"payDebtHistoryW"`
	StatusInHouseParticularCodeW     string `json:"statusInHouseParticularW"`

	IncomePerDebtW             string
	TotalDebtPerYearIncomeW    string
	SavingPerLoanW             string
	MortgageSecuritiesPerLoanW string

	HaveGuarantorCodeW string `json:"haveGuarantorW"`
	IamGuarantorCodeW  string `json:"iamGuarantorW"`
	IncomeTrendCodeW   string `json:"incomeTrendW"`
	LoanObjectCodeW    string `json:"loanObjectW"`
	ProvinceCodeW      string `json:"provinceCodeW"`
}

type questionnaire struct {
	ID        NullInt
	LoanerID  NullInt
	UpdatedBy NullInt
	ApproveBy NullInt

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

type NullString sql.NullString
type NullInt sql.NullInt64
type NullFloat64 sql.NullFloat64
type NullTime sql.NullTime

func (n *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		//n.String, n.Valid = "-", true
		*n = NullString{"-", true}
	} else {
		*n = NullString{s.String, true}
	}
	return nil
}

func (n *NullInt) Scan(value interface{}) error {
	var s sql.NullInt64
	if err := s.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*n = NullInt{0, true}
	} else {
		*n = NullInt{s.Int64, true}
	}
	return nil
}

func (n *NullFloat64) Scan(value interface{}) error {
	var s sql.NullFloat64
	if err := s.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*n = NullFloat64{0, true}
	} else {
		*n = NullFloat64{s.Float64, true}
	}
	return nil
}

func (n *NullTime) Scan(value interface{}) error {
	var s sql.NullTime
	if err := s.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*n = NullTime{time.Time{}, true}
	} else {
		*n = NullTime{s.Time, true}
	}
	return nil
}

func (s *QuestionnaireStruct) checkNumType() error {

	// check is nil
	{
		if s.IncomeInput == nil {
			return ErrIsNilIncomeInput
		}

		if s.LoanInput == nil {
			return ErrIsNilLoanInput
		}

		if s.DebtPerMonthInput == nil {
			return ErrIsNilDebtPerMonthInput
		}

		if s.TotalDebtInput == nil {
			return ErrIsNilTotalDebtInput
		}

		if s.SavingInput == nil {
			return ErrIsNilSavingInput
		}

		if s.MortgageSecuritiesInput == nil {
			return ErrIsNilMortgageSecuritiesInput
		}

	}

	tyIncome := reflect.TypeOf(s.IncomeInput).Kind()
	tyLoan := reflect.TypeOf(s.LoanInput).Kind()
	tyDebtPerMonth := reflect.TypeOf(s.DebtPerMonthInput).Kind()
	tyTotalDebt := reflect.TypeOf(s.TotalDebtInput).Kind()
	tySaving := reflect.TypeOf(s.SavingInput).Kind()
	tyMortgageSecurities := reflect.TypeOf(s.DebtPerMonthInput).Kind()

	// check is string
	{
		if tyIncome == reflect.String {
			return ErrIncomeMustBeNumber
		}

		if tyLoan == reflect.String {
			return ErrLoanMustBeNumber
		}

		if tyDebtPerMonth == reflect.String {
			return ErrDebtPerMonthMustBeNumber
		}

		if tyTotalDebt == reflect.String {
			return ErrTotalDebtMustBeNumber
		}

		if tySaving == reflect.String {
			return ErrSavingMustBeNumber
		}

		if tyMortgageSecurities == reflect.String {
			return ErrMortgageSecuritiesMustBeNumber
		}
	}

	// check is float
	{
		if tyIncome == reflect.Float64 {
			s.Income = reflect.ValueOf(s.IncomeInput).Float()
		}

		if tyLoan == reflect.Float64 {
			s.Loan = reflect.ValueOf(s.LoanInput).Float()
		}

		if tyDebtPerMonth == reflect.Float64 {
			s.DebtPerMonth = reflect.ValueOf(s.DebtPerMonthInput).Float()
		}

		if tyTotalDebt == reflect.Float64 {
			s.TotalDebt = reflect.ValueOf(s.TotalDebtInput).Float()
		}

		if tySaving == reflect.Float64 {
			s.Saving = reflect.ValueOf(s.SavingInput).Float()
		}

		if tyMortgageSecurities == reflect.Float64 {
			s.MortgageSecurities = reflect.ValueOf(s.MortgageSecuritiesInput).Float()
		}
	}

	return nil
}
