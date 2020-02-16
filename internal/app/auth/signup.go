package auth

import (
	"context"
	"github.com/acoshift/pgsql"
	"github.com/asaskevich/govalidator"
	"time"

	//_ "github.com/lib/pq"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/password"
	"strconv"
	"strings"
	"unicode/utf8"
)

type signUpRequest struct {
	CitizenID       string `json:"citizenID"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Phone           string `json:"phone"`
	Birthday        string `json:"birth"` //format 01/01/1990
	GenderID        int    `json:"genderID"`
	MarriedStatusID int    `json:"marriedStatusID"`
	Religion        string `json:"religion"`
	Address         string `json:"address"`
	SubDistrict     string `json:"subDistrict"`
	District        string `json:"district"`
	ProvinceCode    string `json:"provinceCode"`
	Zipcode         string `json:"zipcode"`
}

type signUpResponse struct {
	Message string `json:"message"`
	ID      int64  `json:"id"`
}

// singUp loaner sing up
func signUp(ctx context.Context, req signUpRequest) (res signUpResponse, err error) {
	roleID := 4
	if len(req.CitizenID) != 13 {
		return res, ErrFormatCitizenID
	}

	req.Email = strings.TrimSpace(req.Email)
	req.Email = strings.ToLower(req.Email)

	if req.Email == "" {
		return res, ErrEmailRequired
	}

	if !govalidator.IsEmail(req.Email) {
		return res, ErrEmailInvalid
	}

	if req.Password == "" {
		return res, ErrPasswordRequired
	}
	if n := utf8.RuneCountInString(req.Password); n < 6 || n > 20 {
		return res, ErrPasswordInvalid
	}

	if req.Password != req.PasswordConfirm {
		return res, ErrPasswordNotMatch
	}

	hashedPassword, err := password.Hash(req.Password)
	if err != nil {
		return res, ErrHashPassword
	}

	if r := strings.TrimSpace(req.Religion); r == "" {
		req.Religion = "-"
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Surname = strings.TrimSpace(req.Surname)
	req.Phone = strings.TrimSpace(req.Phone)

	if req.Name == "" {
		return res, ErrNameRequired
	}
	if req.Surname == "" {
		return res, ErrSurNameRequired
	}

	if n := utf8.RuneCountInString(req.Phone); n != 10 {
		return res, ErrPhoneLength
	}
	_, err = strconv.ParseInt(req.Phone, 10, 64)
	if err != nil {
		return res, ErrPhoneMustBeInt
	}

	// check birthday format
	{
		_, err = time.Parse("02/01/2006", req.Birthday)
		if err != nil {
			return res, ErrBirthdayFormat
		}
	}
	/*
		re := regexp.MustCompile(`(0[1-9]|[12]\d|3[01])/(0[1-9]|1[0-2])/([12]\d{3})`)
		if re.MatchString(req.Birthday) == false {
			return res, ErrBirthdayFormat
		}
	*/

	switch req.MarriedStatusID {
	case 1, 2, 3, 4:
	default:
		return res, ErrMarriedStatusIDInvalid
	}

	switch req.GenderID {
	case 1, 2:
	default:
		return res, ErrGenderIDInvalid
	}

	if r := strings.TrimSpace(req.Address); r == "" {
		req.Address = "-"
	}

	if r := strings.TrimSpace(req.SubDistrict); r == "-" || r == "" {
		return res, ErrSubDistrictRequired
	}

	if r := strings.TrimSpace(req.District); r == "-" || r == "" {
		return res, ErrDistrictRequired
	}

	if r := strings.TrimSpace(req.ProvinceCode); r == "" {
		return res, ErrProvinceCodeRequired
	}

	if r := strings.TrimSpace(req.Zipcode); r == "" {
		//return res, ErrZipcodeRequired
	}

	// insert to DB

	{
		child := 0
		facebook := ""
		ig := ""
		line := ""
		addressNull := "-"

		{
			err = dbctx.QueryRow(ctx,
				`insert into users 
				(citizenID, email, password, name, surname, 
				birthday, phone, genderId , marriedId, religion, 
				child, facebook, ig, line, address1_home, 
				address2_home, subDistrict_home, district_home, provinceCode_home, zipCode_home, 
				roleId)
			values
				($1, $2, $3, $4, $5, 
				$6, $7, $8, $9, $10, 
				$11, $12, $13, $14, $15, 
				$16, $17, $18, $19, $20,
				$21)
			returning id
			`, req.CitizenID, req.Email, hashedPassword, req.Name, req.Surname,
				req.Birthday, req.Phone, req.GenderID, req.MarriedStatusID, req.Religion,
				child, facebook, ig, line, req.Address,
				addressNull, req.SubDistrict, req.District, req.ProvinceCode, req.Zipcode,
				roleID).Scan(&res.ID)
		}
		/*
			err = dbctx.QueryRow(ctx,
				`insert into users
					(citizenID, email, password, name, surname,
					birthday, phone, genderId , marriedId, religion,
					child, facebook, ig, line, address1_home,
					address2_home, subDistrict_home, district_home, provinceCode_home, zipCode_home,
					office_name, address1_office, address2_office, subDistrict_office, district_office,
					provinceCode_office, zipCode_office, roleId)
				values
					($1, $2, $3, $4, $5,
					$6, $7, $8, $9, $10,
					$11, $12, $13, $14, $15,
					$16, $17, $18, $19, $20,
					$21, $22, $23, $24, $25,
					$26, $27, $28)
				returning id
			`, req.CitizenID, req.Email, hashedPassword, req.Name, req.Surname,
				req.Birthday, req.Phone, req.GenderID, req.MarriedStatusID, req.Religion,
				child, facebook, ig, line, req.Address,
				addressNull, req.SubDistrict, req.District, req.ProvinceCode, req.Zipcode,
				"-", req.Address, addressNull, req.SubDistrict, req.District,
				req.ProvinceCode, req.Zipcode, roleID).Scan(&res.ID)

		*/

		if pgsql.IsUniqueViolation(err, "users_email_idx") {
			return res, ErrEmailDuplicated
		}

		if pgsql.IsUniqueViolation(err, "users_phone_idx") {
			return res, ErrPhoneDuplicated
		}

		if pgsql.IsForeignKeyViolation(err, "users_provincecode_home_fkey") {
			return res, ErrProvinceCodeInvalid
		}

		if err != nil {
			return res, err
		}
	}
	res.Message = "การลงทะเบียนสำเร็จ"
	return
}
