package admin

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/password"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/user"
	"strconv"
	"strings"
	"unicode/utf8"
)

type createNewWorkerRequest struct {
	Name            string `json:"name"`
	Surname         string `json:"surname"`
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
	userRole := auth.GetUserRole(ctx)

	if userRole != 1 {
		return res, ErrPermissionNotAllow
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

	hashedPassword, err := password.Hash(req.Password)
	if err != nil {
		return res, ErrHashPassword
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

	arg := user.SignUpArgs{
		Name:            req.Name,
		Surname:         req.Surname,
		Email:           req.Email,
		Phone:           req.Phone,
		HashPassword:    hashedPassword,
		Birthday:        req.Birthday,
		GenderID:        req.GenderID,
		MarriedStatusID: req.MarriedStatusID,
		Religion:        req.Religion,
		RoleId:          req.RoleID}

	_, err = user.Insert(ctx, &arg)
	if err == user.ErrEmailDuplicated {
		return res, ErrEmailNotAvailable
	}

	if err == user.ErrPhoneDuplicated {
		return res, ErrPhoneNotAvailable
	}
	if err != nil {
		fmt.Println(err)
		return res, ErrSomething
	}
	res.Message = "สร้างบัญชีพนักงานสำเร็จ"

	return

}
