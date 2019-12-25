package questionnaireOption

import (
	"errors"
	"net/http"
)

var (
	ErrGetOptionAge = errors.New("get age option")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrGetOptionAge:
		return http.StatusInternalServerError

	default:
		return http.StatusInternalServerError
	}
}

func errorToMessage(err error) string {
	switch err {
	case ErrGetOptionAge:
		return "เกิดข้อผิดพลาดให้การเรียกข้อมูลตัวเลือกอายุ"
	default:
		return "internal server error"

	}
}
