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

var (
	ErrAddress1Required     = errors.New("updateHomeAddress: address1 invalid")
	ErrAddress2Required     = errors.New("updateHomeAddress: address2 invalid")
	ErrSubDistrictRequired  = errors.New("updateHomeAddress: sub district invalid")
	ErrDistrictRequired     = errors.New("updateHomeAddress: district invalid")
	ErrProvinceCodeRequired = errors.New("updateHomeAddress: province code invalid")
	ErrZipcodeRequired      = errors.New("updateHomeAddress: zipCode invalid")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrUserNotLogin:
		return http.StatusNotAcceptable
	case ErrNameRequired, ErrSurnameRequired, ErrPhoneRequired, ErrBirthRequired:
		return http.StatusBadRequest
	case ErrAddress1Required, ErrAddress2Required, ErrSubDistrictRequired, ErrDistrictRequired, ErrProvinceCodeRequired, ErrZipcodeRequired:
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

	case ErrAddress1Required:
		return "กรุณาระบุที่อยู่"
	case ErrAddress2Required:
		return ""
	case ErrSubDistrictRequired:
		return "กรุณาเลือกตำบลที่อยู่ของท่าน"
	case ErrDistrictRequired:
		return "กรุณาเลือกอำเภอที่อยู่ของท่าน"
	case ErrProvinceCodeRequired:
		return "กรุณาเลือกจังหวัดที่อยู่ของท่าน"
	case ErrZipcodeRequired:
		return "กรุณาระบุรหัสไปรษณีที่อยู่ของท่าน"
	default:
		return "internal server error"
	}
}
