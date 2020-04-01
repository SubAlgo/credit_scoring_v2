package questionnaire

import (
	"context"
	"math"
)

type prepareArgs struct {
	/* เป้าหมายคำนวณค่าจาก input
	เพื่อให้ได้ creditGrade, creditRisk, riskLevel, matrixIndex

	*/
	SuggestScore     int
	SuggestGiveScore int

	income             float64
	loan               float64
	debtPerMonth       float64
	totalDebt          float64
	saving             float64
	mortgageSecurities float64

	incomePerDebtCode             string
	totalDebtPerYearIncomeCode    string
	savingPerLoanCode             string
	mortgageSecuritiesPerLoanCode string

	ageCode                         string // code option
	jobCode                         string
	eduCode                         string
	timeJobCode                     string
	freChangeNameCode               string
	timeOfPhoneNumberCode           string
	timeOfNameInHouseParticularCode string
	payDebtHistoryCode              string
	statusInHouseParticularCode     string

	haveGuarantorCode string
	iamGuarantorCode  string
	incomeTrendCode   string
	loanObjectCode    string
	provinceCode      string

	creditGrade string
	creditRisk  string
	riskLevel   string
	matrixIndex string
}

func (p *prepareArgs) prepareData(ctx context.Context) (err error) {
	/*
		จัดเตรียมข้อมูล,​เช็คค่าว่าง
		และ คำนวณค่าของ
			IncomePerDebt,
			TotalDebtPerYearIncome,
			SavingPerLoan,
			MortgageSecuritiesPerLoan
	*/
	// check input
	{
		if p.ageCode == "" {
			return ErrInvalidInputAge
		}

		if p.income <= 0 {
			return ErrInvalidInputIncome
		}

		if p.loan <= 0 {
			return ErrInvalidInputLoan
		}

		if p.debtPerMonth < 0 {
			return ErrInvalidInputDebtPerMonth
		}

		if p.totalDebt < 0 {
			return ErrInvalidInputTotalDebt
		}

		if p.saving < 0 {
			return ErrInvalidInputSaving
		}

		if p.mortgageSecurities < 0 {
			return ErrInvalidInputMortgageSecurities
		}

		if p.ageCode == "" || p.ageCode == "0" {
			return ErrInvalidInputAge
		}

		/*
			if p.ageCode == "6" {
				return ErrYourAgeCanNotLowerThan20
			}
		*/

		if p.jobCode == "" || p.jobCode == "0" {
			return ErrInvalidInputJob
		}

		if p.eduCode == "" || p.eduCode == "0" {
			return ErrInvalidInputEdu
		}

		if p.timeJobCode == "" || p.timeJobCode == "0" {
			return ErrInvalidInputTimeJob
		}

		if p.freChangeNameCode == "" || p.freChangeNameCode == "0" {
			return ErrInvalidInputFreChangeName
		}

		if p.timeOfPhoneNumberCode == "" || p.timeOfPhoneNumberCode == "0" {
			return ErrInvalidInputTimeOfPhoneNumber
		}

		if p.timeOfNameInHouseParticularCode == "" || p.timeOfNameInHouseParticularCode == "0" {
			return ErrInvalidInputTimeOfStayInHouseParticular
		}

		if p.payDebtHistoryCode == "" || p.payDebtHistoryCode == "0" {
			return ErrInvalidInputPayDebtHistory
		}

		if p.statusInHouseParticularCode == "" || p.statusInHouseParticularCode == "0" {
			return ErrInvalidInputStatusInHouseParticular
		}

		if p.haveGuarantorCode == "" || p.haveGuarantorCode == "0" {
			return ErrInvalidInputHaveGuarantor
		}

		if p.iamGuarantorCode == "" || p.iamGuarantorCode == "0" {
			return ErrInvalidInputIamGuarantor
		}

		if p.incomeTrendCode == "" || p.incomeTrendCode == "0" {
			return ErrInvalidInputIncomeTrend
		}

		if p.loanObjectCode == "" || p.loanObjectCode == "0" {
			return ErrInvalidInputLoanObject
		}

		if p.provinceCode == "" || p.provinceCode == "0" {
			return ErrInvalidInputProvinceCode
		}
	}

	// cal code incomePerDebt, totalDebtPerYearIncome ,savingPerLoan, mortgageSecuritiesPerLoan
	{
		incomePerDebt := math.Round(((p.income - p.debtPerMonth) / p.debtPerMonth) * 100)
		totalDebtPerYearIncome := math.Round((p.totalDebt / (p.income * 12)) * 100)
		savingPerLoan := math.Round((p.saving / p.loan) * 100)
		{
			if incomePerDebt > 80 {
				p.incomePerDebtCode = "5"
			} else if incomePerDebt >= 51 && incomePerDebt <= 80 {
				p.incomePerDebtCode = "4"
			} else if incomePerDebt >= 21 && incomePerDebt <= 50 {
				p.incomePerDebtCode = "3"
			} else if incomePerDebt >= 1 && incomePerDebt <= 20 {
				p.incomePerDebtCode = "2"
			} else {
				p.incomePerDebtCode = "1"
			}
		}

		{
			if totalDebtPerYearIncome <= 29 {
				p.totalDebtPerYearIncomeCode = "5"
			} else if totalDebtPerYearIncome >= 30 && totalDebtPerYearIncome <= 39 {
				p.totalDebtPerYearIncomeCode = "4"
			} else if totalDebtPerYearIncome < 40 {
				p.totalDebtPerYearIncomeCode = "3"
			} else if totalDebtPerYearIncome == 40 {
				p.totalDebtPerYearIncomeCode = "2"
			} else if totalDebtPerYearIncome > 40 {
				p.totalDebtPerYearIncomeCode = "1"
			}
		}

		{
			if savingPerLoan > 20 {
				p.savingPerLoanCode = "5"
			} else if savingPerLoan >= 15 && savingPerLoan <= 20 {
				p.savingPerLoanCode = "4"
			} else if savingPerLoan >= 10 && savingPerLoan <= 14 {
				p.savingPerLoanCode = "3"
			} else if savingPerLoan >= 5 && savingPerLoan <= 9 {
				p.savingPerLoanCode = "2"
			} else if savingPerLoan >= 0 && savingPerLoan <= 4 { // 0 - 4
				p.savingPerLoanCode = "1"
			}
		}

		{
			if p.mortgageSecurities > p.loan {
				p.mortgageSecuritiesPerLoanCode = "5"
			} else if p.mortgageSecurities == p.loan {
				p.mortgageSecuritiesPerLoanCode = "4"
			} else if p.mortgageSecurities < p.loan {
				p.mortgageSecuritiesPerLoanCode = "1"
			}
		}
	}

	// get quiz score data
	var (
		ageScore                         int
		jobScore                         int
		eduScore                         int
		timeJobScore                     int
		freChangeNameScore               int
		timeOfPhoneNumberScore           int
		timeOfNameInHouseParticularScore int
		payDebtHistoryScore              int
		statusInHouseParticularScore     int

		incomePerDebtScore          int
		totalDebtPerYearIncomeScore int

		savingPerLoanScore int

		mortgageSecuritiesPerLoanScore int
		haveGuarantorScore             int
		iamGuarantorScore              int

		incomeTrendScore int
		loanObjectScore  int
		provinceScore    int
	)

	// get quiz score
	{
		ageScore, err = getAgeScore(ctx, p.ageCode)
		if err != nil {
			return ErrGetAgeScoreFromDB
		}
		jobScore, err = getJobScore(ctx, p.jobCode)
		if err != nil {
			return ErrGetJobScoreFromDB
		}
		eduScore, err = getEduScore(ctx, p.eduCode)
		if err != nil {
			return ErrGetEduScoreFromDB
		}

		timeJobScore, err = getTimeJobScore(ctx, p.timeJobCode)
		if err != nil {
			return ErrGetTimeJobScoreFromDB
		}

		freChangeNameScore, err = getFreChangeNameScore(ctx, p.freChangeNameCode)
		if err != nil {
			return ErrGetFreChangeNameScoreFromDB
		}

		timeOfPhoneNumberScore, err = getTimeOfPhoneNumberScore(ctx, p.timeOfPhoneNumberCode)
		if err != nil {
			return ErrGetAgeScoreFromDB
		}

		timeOfNameInHouseParticularScore, err = getTimeOfNameInHouseParticularScore(ctx, p.timeOfNameInHouseParticularCode)
		if err != nil {
			return ErrGetTimeOfNameInHouseParticularScoreFromDB
		}

		payDebtHistoryScore, err = getPayDebtHistoryScore(ctx, p.payDebtHistoryCode)
		if err != nil {
			return ErrGetPayDebtHistoryScoreFromDB
		}

		statusInHouseParticularScore, err = getStatusInHouseParticularScore(ctx, p.statusInHouseParticularCode)
		if err != nil {
			return ErrGetStatusInHouseParticularScoreFromDB
		}

		incomePerDebtScore, err = getIncomePerDebtScore(ctx, p.incomePerDebtCode)
		if err != nil {
			return ErrGetIncomePerDebtScoreFromDB
		}

		totalDebtPerYearIncomeScore, err = getTotalDebtPerYearIncomeScore(ctx, p.totalDebtPerYearIncomeCode)
		if err != nil {
			return ErrGetTotalDebtPerYearIncomeScoreFromDB
		}

		savingPerLoanScore, err = getSavingPerLoanScore(ctx, p.savingPerLoanCode)
		if err != nil {
			return ErrGetSavingPerLoanScoreFromDB
		}

		mortgageSecuritiesPerLoanScore, err = getMortgageSecuritiesPerLoanScore(ctx, p.mortgageSecuritiesPerLoanCode)
		if err != nil {
			return ErrGetMortgageSecuritiesPerLoanScoreFromDB
		}

		haveGuarantorScore, err = getHaveGuarantorScore(ctx, p.haveGuarantorCode)
		if err != nil {
			return ErrGetHaveGuarantorScoreFromDB
		}

		iamGuarantorScore, err = getIamGuarantorScore(ctx, p.iamGuarantorCode)
		if err != nil {
			return ErrGetIamGuarantorScoreFromDB
		}

		incomeTrendScore, err = getIncomeTrendScore(ctx, p.incomeTrendCode)
		if err != nil {
			return ErrGetIncomeTrendScoreFromDB
		}

		loanObjectScore, err = getLoanObjectScore(ctx, p.loanObjectCode)
		if err != nil {
			return ErrGetLoanObjectScoreFromDB
		}

		provinceScore, err = getProvinceScore(ctx, p.provinceCode)
		if err != nil {
			return ErrGetProvinceScoreFromDB
		}
	}
	var (
		criteriaScore1     float64
		criteriaScore2     float64
		criteriaScore3     float64
		criteriaScore4     float64
		criteriaScore5     float64
		totalCriteriaScore float64
	)

	// calculate criteria score
	{
		if p.SuggestScore == 0 {
			criteriaScore1 = 3 * calCriteriaScore([]int{ageScore, jobScore, eduScore, timeJobScore, freChangeNameScore, timeOfPhoneNumberScore, timeOfNameInHouseParticularScore, payDebtHistoryScore, statusInHouseParticularScore})
		} else {
			criteriaScore1 = 3 * calCriteriaScore([]int{p.SuggestScore, p.SuggestGiveScore, ageScore, jobScore, eduScore, timeJobScore, freChangeNameScore, timeOfPhoneNumberScore, timeOfNameInHouseParticularScore, payDebtHistoryScore, statusInHouseParticularScore})
		}

		criteriaScore2 = 3 * calCriteriaScore([]int{incomePerDebtScore, totalDebtPerYearIncomeScore})

		criteriaScore3 = 2 * calCriteriaScore([]int{savingPerLoanScore})

		criteriaScore4 = 1 * calCriteriaScore([]int{mortgageSecuritiesPerLoanScore, haveGuarantorScore, iamGuarantorScore})

		criteriaScore5 = 1 * calCriteriaScore([]int{incomeTrendScore, loanObjectScore, provinceScore})

		totalCriteriaScore = criteriaScore1 + criteriaScore2 + criteriaScore3 + criteriaScore4 + criteriaScore5
	}

	//cal credit grade
	{
		if totalCriteriaScore >= 800 {
			p.creditGrade = "A"
		} else if totalCriteriaScore >= 700 && totalCriteriaScore < 800 {
			p.creditGrade = "B"
		} else if totalCriteriaScore >= 600 && totalCriteriaScore < 700 {
			p.creditGrade = "C"
		} else if totalCriteriaScore >= 500 && totalCriteriaScore < 600 {
			p.creditGrade = "D"
		} else if totalCriteriaScore < 500 {
			p.creditGrade = "F"
		}
	}
	/*
		{
			if totalCriteriaScore >= 800 && totalCriteriaScore <= 1000 {
				p.creditGrade = "A"
			} else if totalCriteriaScore >= 700 && totalCriteriaScore <= 799 {
				p.creditGrade = "B"
			} else if totalCriteriaScore >= 600 && totalCriteriaScore <= 699 {
				p.creditGrade = "C"
			} else if totalCriteriaScore >= 500 && totalCriteriaScore <= 599 {
				p.creditGrade = "D"
			} else if totalCriteriaScore <= 499 {
				p.creditGrade = "F"
			}
		}
	*/

	//cal credit risk
	{
		criteria1Percent := (criteriaScore1 / 300) * 100
		criteria2Percent := (criteriaScore2 / 300) * 100
		criteria3Percent := (criteriaScore3 / 200) * 100

		characterGrade := criteriaPercentToGrade(criteria1Percent)
		capacityGrade := criteriaPercentToGrade(criteria2Percent)
		capitalGrade := criteriaPercentToGrade(criteria3Percent)

		if (characterGrade == "A" || characterGrade == "B") && (capacityGrade == "A" || capacityGrade == "B") && (capitalGrade == "A" || capitalGrade == "B") {
			p.creditRisk = "1"
		} else if (characterGrade == "A" || characterGrade == "B") && (capacityGrade == "A" || capacityGrade == "B") && (capitalGrade == "C" || capitalGrade == "D" || capitalGrade == "F") {
			p.creditRisk = "2"
		} else if (characterGrade == "A" || characterGrade == "B") && (capacityGrade == "C" || capacityGrade == "D" || capacityGrade == "F") && (capitalGrade == "A" || capitalGrade == "B") {
			p.creditRisk = "3"
		} else if (characterGrade == "A" || characterGrade == "B") && (capacityGrade == "C" || capacityGrade == "D" || capacityGrade == "F") && (capitalGrade == "C" || capitalGrade == "D" || capitalGrade == "F") {
			p.creditRisk = "3"
		} else if (characterGrade == "C" || characterGrade == "D" || characterGrade == "F") && (capacityGrade == "A" || capacityGrade == "B") && (capitalGrade == "A" || capitalGrade == "B") {
			p.creditRisk = "4"
		} else if (characterGrade == "C" || characterGrade == "D" || characterGrade == "F") && (capacityGrade == "C" || capacityGrade == "D" || capacityGrade == "F") && (capitalGrade == "A" || capitalGrade == "B") {
			p.creditRisk = "4"
		} else if (characterGrade == "C" || characterGrade == "D" || characterGrade == "F") && (capacityGrade == "B" || capacityGrade == "A") && (capitalGrade == "C" || capitalGrade == "D" || capitalGrade == "F") {
			p.creditRisk = "5"
		} else {
			p.creditRisk = "5"
		}
	}

	//set matrixIndex
	{
		p.matrixIndex = p.creditGrade + p.creditRisk
	}

	//set result risk level
	{
		switch p.matrixIndex {
		case "F5", "F4", "D5":
			p.riskLevel = "Very High Risk"
		case "C5", "D4", "F3":
			p.riskLevel = "High Risk"
		case "A3", "B2", "C1":
			p.riskLevel = "Low Risk"
		case "A2", "A1", "B1":
			p.riskLevel = "Very Low Risk"
		case "A5", "A4", "B5", "B4", "B3", "C4", "C3", "C2", "D3", "D2", "D1", "F2", "F1":
			p.riskLevel = "Medium Risk"
		}
	}
	return
}

func calCriteriaScore(s []int) (score float64) {
	n := len(s)
	maxScore := float64(100 / n)
	minScore := float64(maxScore / 5)
	score2 := float64(minScore * 2)
	score3 := float64(minScore * 3)
	score4 := float64(minScore * 4)
	score = 0

	for _, qScore := range s {
		switch qScore {
		case 1:
			score = score + minScore
		case 2:
			score = score + score2
		case 3:
			score = score + score3
		case 4:
			score = score + score4
		case 5:
			score = score + maxScore
		}
	}
	return
}

func criteriaPercentToGrade(criPercent float64) (criteriaGrade string) {
	if criPercent >= 80 && criPercent <= 100 {
		criteriaGrade = "A"
	} else if criPercent >= 70 && criPercent <= 79 {
		criteriaGrade = "B"
	} else if criPercent >= 60 && criPercent <= 69 {
		criteriaGrade = "C"
	} else if criPercent >= 50 && criPercent <= 59 {
		criteriaGrade = "D"
	} else if criPercent < 50 {
		criteriaGrade = "F"
	}
	return
}
