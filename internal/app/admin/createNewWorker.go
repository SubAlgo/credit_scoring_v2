package admin

import (
	"context"
	"fmt"
	"github.com/acoshift/pgsql"
	"github.com/asaskevich/govalidator"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/password"
	"strconv"
	"strings"
	"unicode/utf8"
)

type createNewWorkerRequest struct {
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	CitizenID       string `json:"citizenID"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	Birthday        string `json:"birth"`
	GenderID        int    `json:"genderID"`
	MarriedStatusID int    `json:"marriedStatusID"`
	Religion        string `json:"religion"`
	RoleID          int    `json:"roleID"`
}

type processResponse struct {
	Message string `json:"message"`
}

func createNewWorker(ctx context.Context, req createNewWorkerRequest) (res processResponse, err error) {

	req.MarriedStatusID = 1

	// check role is superAdmin
	{
		userRole := auth.GetUserRole(ctx)
		if userRole != 1 {
			return res, ErrPermissionNotAllow
		}
	}

	// Check input (name)
	{
		req.Name = strings.TrimSpace(req.Name)
		if req.Name == "" {
			return res, ErrNameRequired
		}
	}

	// Check input (surname)
	{
		req.Surname = strings.TrimSpace(req.Surname)
		if req.Surname == "" {
			return res, ErrSurNameRequired
		}
	}

	// Check input (citizen id)
	{
		req.CitizenID = strings.TrimSpace(req.CitizenID)
		if req.CitizenID == "" {
			return res, ErrCitizenIDRequired
		}

		if len(req.CitizenID) != 13 {
			return res, ErrCitizenIDInvalid
		}
	}

	{
		req.Email = strings.TrimSpace(req.Email)
		req.Email = strings.ToLower(req.Email)

		if !govalidator.IsEmail(req.Email) {
			return res, ErrEmailInvalid
		}

		if req.Email == "" {
			return res, ErrEmailRequired
		}
	}

	{
		req.Phone = strings.TrimSpace(req.Phone)
		if n := utf8.RuneCountInString(req.Phone); n != 10 {
			return res, ErrPhoneLength
		}
		_, err = strconv.ParseInt(req.Phone, 10, 64)
		if err != nil {
			return res, ErrPhoneMustBeInt
		}
	}

	if req.Password == "" {
		return res, ErrPasswordRequired
	}
	if n := utf8.RuneCountInString(req.Password); n < 6 || n > 20 {
		return res, ErrPasswordInvalid
	}

	hashedPassword, err := password.Hash(req.Password)
	if err != nil {
		return res, ErrHashPassword
	}

	_, err = req.insert(ctx, hashedPassword)

	if pgsql.IsUniqueViolation(err, "users_citizenid_idx") {
		return res, ErrCitizenIDDuplicated
	}

	if err != nil {
		fmt.Println(err)
		return res, ErrSomething
	}

	res.Message = "สร้างบัญชีพนักงานสำเร็จ"

	return

}
