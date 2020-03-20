package forgotPassword

import (
	"errors"
	"net/http"
)

var (
	ErrGetForgotPasswordQuestion = errors.New("forgot password: get question")
	ErrEmailRequired             = errors.New("forgot password: email required")
	ErrNoResult                  = errors.New("forgot password: no result")
	ErrGetQuestionByEmail        = errors.New("forgot password: get question by email")
	ErrAnswerRequired            = errors.New("forgot password: answer required")
	ErrEmailAndAnswerNotMarch    = errors.New("forgot password: email and answer not match")
	ErrUpdateNewPassword         = errors.New("forgot password: update new password")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrGetForgotPasswordQuestion:
		return http.StatusInternalServerError
	case ErrEmailRequired, ErrNoResult, ErrGetQuestionByEmail, ErrAnswerRequired, ErrEmailAndAnswerNotMarch:
		return http.StatusBadRequest
	case ErrUpdateNewPassword:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func errorToMessage(err error) string {
	switch err {
	case ErrGetForgotPasswordQuestion:
		return "internal server error (get forgot password question)"
	case ErrEmailRequired:
		return "กรุณาระบุอีเมล์ของท่าน"
	case ErrNoResult, ErrGetQuestionByEmail:
		return "ไม่พบข้อมูลอาจเนื่องจากท่านได้ระบุอีเมล์ไม่ถูกต้อง"
	case ErrAnswerRequired:
		return "กรุณาระบุคำตอบสำหรับกู้รหัสผ่าน"
	case ErrEmailAndAnswerNotMarch:
		return "ระบุคำตอบไม่ถูกต้อง"
	case ErrUpdateNewPassword:
		return "internal server error (update new password)"
	default:
		return "internal server error"
	}
}
