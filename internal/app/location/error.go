package location

import (
	"errors"
	"net/http"
)

var (
	ErrMakeSubDistrictList = errors.New("location: make sub-district list")
)

func errorToStatusCode(err error) int {
	switch err {
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
	default:
		return "internal server error"

	}
}
