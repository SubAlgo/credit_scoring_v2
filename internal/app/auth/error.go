package auth

import (
	"errors"
	"net/http"
)

var (
	ErrEmailRequired      = errors.New("auth: email required")
	ErrPasswordRequires   = errors.New("auth: password required")
	ErrInvalidCredentials = errors.New("auth: invalid credentials")
	ErrEmailInvalid       = errors.New("auth: email invalid")
	ErrPasswordInvalid    = errors.New("auth: password invalid")
	ErrPasswordRequired   = errors.New("auth: password required")
	ErrHashPassword       = errors.New("auth: hash password has problem")
	ErrNameRequired       = errors.New("auth: name required")
	ErrSurNameRequired    = errors.New("auth: surname required")
	ErrPhoneLength        = errors.New("auth: phone number length not equal 10")
	ErrPhoneMustBeInt     = errors.New("auth: phone number must be integer")
	ErrEmailNotAvailable  = errors.New("auth: email not available")
	ErrPhoneNotAvailable  = errors.New("auth: phone not available")
	ErrSomething          = errors.New("auth: error something")
	ErrTokenRequired      = errors.New("auth: token required")
	ErrNotSignIn          = errors.New("auth: not sign in")

	ErrUsernameRequired = errors.New("auth: required email or phone number")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrEmailRequired, ErrEmailInvalid, ErrPasswordRequires, ErrPasswordInvalid, ErrNameRequired, ErrSurNameRequired, ErrPhoneLength, ErrPhoneMustBeInt, ErrEmailNotAvailable, ErrPhoneNotAvailable:
		return http.StatusBadRequest

	case ErrUsernameRequired, ErrTokenRequired:
		return http.StatusBadRequest
	case ErrInvalidCredentials, ErrNotSignIn:
		return http.StatusUnauthorized
	case ErrHashPassword, ErrSomething:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func errorToMessage(err error) string {
	switch err {
	case ErrEmailInvalid:
		return "รูปแบบอีเมลไม่ถูกต้อง"
	case ErrPasswordRequires:
		return "กรุณากำหนดระหัสผ่าน"
	case ErrPasswordInvalid:
		return "พาสเวิร์ดต้องมีความยาว 6 - 20 ตัวอักษร"

	case ErrNameRequired:
		return "กรุณาระบุชื่อของท่าน"
	case ErrSurNameRequired:
		return "กรุณาระบุนามสกุลของท่าน"
	case ErrPhoneLength:
		return "เบอร์โทรต้องมีความยาว 10 ตัว"
	case ErrPhoneMustBeInt:
		return "เบอร์โทรต้องเป็นตัวเลขเท่านั้้น"
	case ErrEmailNotAvailable:
		return "อีเมลนี้ได้ใช้ทำการลงทะเบียนเรียบร้อยแล้ว ท่านไม่สามารถใช้อีเมลนีั้ได้"
	case ErrPhoneNotAvailable:
		return "เบอร์โทรนี้ได้ใช้ทำการลงทะเบียนเรียบร้อยแล้ว ท่านไม่สามารถใช้เบอร์โทรนีั้ได้"

	case ErrEmailRequired:
		return "กรุณาระบุ อีเมล์ หรือ เบอร์โทรศัพท์ เพื่อเข้าสู่ระบบ"
	case ErrInvalidCredentials:
		return "อีเมลหรือหมายเลขโทรศัพท์ที่คุณป้อนไม่ตรงกับบัญชีผู้ใช้ใดๆ"
	case ErrTokenRequired:
		return "token required"
	case ErrNotSignIn:
		return "ท่านยังไม่ได้เข้าสู่ระบบ"

	case ErrSomething:
		return "Internal server error"
	default:
		return "Internal Server Error"
	}

}
