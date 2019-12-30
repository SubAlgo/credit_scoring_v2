package province

import (
	"errors"
	"net/http"
)

var (
	ErrUserNotLogin         = errors.New("updateProfile: signIn required")
	ErrInvalidProvinceScore = errors.New("setProvinceScore: error invalid province score")
	ErrInvalidProvinceID    = errors.New("setProvinceScore: error invalid province id")
	ErrSelectProvinceList   = errors.New("select province list had error")
	ErrCreateProvinceList   = errors.New("create province list had error")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrUserNotLogin, ErrInvalidProvinceScore, ErrInvalidProvinceID:
		return http.StatusBadRequest
	case ErrSelectProvinceList, ErrCreateProvinceList:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}

func errorToMessage(err error) string {
	switch err {
	case ErrUserNotLogin:
		return "ท่านยังไม่ได้ login เข้าสู่ระบบ"
	case ErrInvalidProvinceID:
		return "กรุณาระบุรหัสจังหวัด"
	case ErrInvalidProvinceScore:
		return "กรุณาระบุคะแนนของจังหวัด"

	case ErrSelectProvinceList, ErrCreateProvinceList:
		return "Internal Server Error"
	default:
		return "Internal Server Error"
	}
}
