package auth

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/password"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/user"
	"strconv"
	"strings"
	"unicode/utf8"
)

type signUpRequest struct {
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	Birthday        string `json:"birth"`
	GenderID        int    `json:"genderID"`
	MarriedStatusID int    `json:"marriedStatusID"`
	Religion        string `json:"religion"`
}

type signUpResponse struct {
	ID int64 `json:"id"`
}

// singUp loaner sing up 
func signUp(ctx context.Context, req signUpRequest) (res signUpResponse, err error) {
	roleID := 4

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
		RoleId:          roleID}

	res.ID, err = user.Insert(ctx, &arg)
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

	return

}
