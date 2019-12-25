package location

import (
	"errors"
	"net/http"
)

var (
	ErrMakeSubDistrictList = errors.New("location: make sub-district list")
	ErrZipcodeInvalid      = errors.New("location: zipcode invalid")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrZipcodeInvalid:
		return http.StatusBadRequest
	case ErrMakeSubDistrictList:
		return http.StatusInternalServerError

	default:
		return http.StatusInternalServerError
	}
}

func errorToMessage(err error) string {
	switch err {
	case ErrMakeSubDistrictList:
		return "เกิดข้อผิดพลาดในการดึงข้อมูลตำบล"
	case ErrZipcodeInvalid:
		return "ระบุรหัสไปรษณีย์ไม่ถูกต้อง"
	default:
		return "internal server error"

	}
}
