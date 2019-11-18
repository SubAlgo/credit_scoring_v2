package user

import (
	"errors"
	"net/http"
)

var (
	ErrUserNotLogin = errors.New("user: not login")

	ErrUpdateProfile = errors.New("updateProfile: update profile error")
)

var (
	ErrNameRequired    = errors.New("updateProfile: name invalid")
	ErrSurnameRequired = errors.New("updateProfile: surname invalid")
	ErrPhoneRequired   = errors.New("updateProfile: phone invalid")
	ErrBirthRequired   = errors.New("updateProfile: name invalid")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrUserNotLogin:
		return http.StatusNotAcceptable
	case ErrNameRequired, ErrSurnameRequired, ErrPhoneRequired, ErrBirthRequired:
		return http.StatusBadRequest
	case ErrUpdateProfile:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func errorToMessage(err error) string {
	switch err {
	case ErrUserNotLogin:
		return "ท่านยังไม่ได้เข้าสู่ระบบ"
	default:
		return "internal server error"
	}
}
