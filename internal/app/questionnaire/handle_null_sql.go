package questionnaire

import (
	"database/sql"
	"reflect"
	"time"
)

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

func (s *QuestionnaireStruct) checkNumTypeForWorker() error {

	// check is nil
	{
		if s.IncomeInputW == nil {
			return ErrIsNilIncomeInput
		}

		if s.LoanInputW == nil {
			return ErrIsNilLoanInput
		}

		if s.DebtPerMonthInputW == nil {
			return ErrIsNilDebtPerMonthInput
		}

		if s.TotalDebtInputW == nil {
			return ErrIsNilTotalDebtInput
		}

		if s.SavingInputW == nil {
			return ErrIsNilSavingInput
		}

		if s.MortgageSecuritiesInputW == nil {
			return ErrIsNilMortgageSecuritiesInput
		}

	}

	tyIncome := reflect.TypeOf(s.IncomeInputW).Kind()
	tyLoan := reflect.TypeOf(s.LoanInputW).Kind()
	tyDebtPerMonth := reflect.TypeOf(s.DebtPerMonthInputW).Kind()
	tyTotalDebt := reflect.TypeOf(s.TotalDebtInputW).Kind()
	tySaving := reflect.TypeOf(s.SavingInputW).Kind()
	tyMortgageSecurities := reflect.TypeOf(s.DebtPerMonthInputW).Kind()

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
			s.IncomeW = reflect.ValueOf(s.IncomeInputW).Float()
		}

		if tyLoan == reflect.Float64 {
			s.LoanW = reflect.ValueOf(s.LoanInputW).Float()
		}

		if tyDebtPerMonth == reflect.Float64 {
			s.DebtPerMonthW = reflect.ValueOf(s.DebtPerMonthInputW).Float()
		}

		if tyTotalDebt == reflect.Float64 {
			s.TotalDebtW = reflect.ValueOf(s.TotalDebtInputW).Float()
		}

		if tySaving == reflect.Float64 {
			s.SavingW = reflect.ValueOf(s.SavingInputW).Float()
		}

		if tyMortgageSecurities == reflect.Float64 {
			s.MortgageSecuritiesW = reflect.ValueOf(s.MortgageSecuritiesInputW).Float()
		}
	}

	return nil
}
