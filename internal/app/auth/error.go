package auth

import (
	"errors"
	"net/http"
)

var (
	ErrEmailRequired          = errors.New("auth: email required")
	ErrPasswordRequires       = errors.New("auth: password required")
	ErrPasswordNotMatch       = errors.New("auth: password not match")
	ErrInvalidCredentials     = errors.New("auth: invalid credentials")
	ErrEmailInvalid           = errors.New("auth: email invalid")
	ErrPasswordInvalid        = errors.New("auth: password invalid")
	ErrPasswordRequired       = errors.New("auth: password required")
	ErrHashPassword           = errors.New("auth: hash password has problem")
	ErrNameRequired           = errors.New("auth: name required")
	ErrSurNameRequired        = errors.New("auth: surname required")
	ErrPhoneLength            = errors.New("auth: phone number length not equal 10")
	ErrPhoneMustBeInt         = errors.New("auth: phone number must be integer")
	ErrEmailNotAvailable      = errors.New("auth: email not available")
	ErrEmailDuplicated        = errors.New("auth: email duplicated")
	ErrPhoneDuplicated        = errors.New("auth: phone duplicated")
	ErrPhoneNotAvailable      = errors.New("auth: phone not available")
	ErrSomething              = errors.New("auth: error something")
	ErrTokenRequired          = errors.New("auth: token required")
	ErrNotSignIn              = errors.New("auth: not sign in")
	ErrSubDistrictRequired    = errors.New("auth: subDistrict required")
	ErrDistrictRequired       = errors.New("auth: district required")
	ErrProvinceCodeRequired   = errors.New("auth: province required")
	ErrGenderIDInvalid        = errors.New("auth: gender id invalid")
	ErrMarriedStatusIDInvalid = errors.New("auth: married status id invalid")
	ErrZipcodeRequired        = errors.New("auth: zipcode required")
	ErrProvinceCodeInvalid    = errors.New("auth: province code invalid")
	ErrUsernameRequired       = errors.New("auth: required email or phone number")
	ErrBirthdayFormat         = errors.New("auth: birthday format error")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrEmailRequired, ErrEmailInvalid, ErrPasswordRequires, ErrPasswordInvalid, ErrNameRequired, ErrSurNameRequired, ErrPhoneLength, ErrPhoneMustBeInt, ErrEmailNotAvailable, ErrPhoneNotAvailable:
		return http.StatusBadRequest

	case ErrUsernameRequired, ErrTokenRequired, ErrPasswordNotMatch, ErrEmailDuplicated, ErrPhoneDuplicated, ErrSubDistrictRequired, ErrDistrictRequired, ErrProvinceCodeRequired, ErrGenderIDInvalid:
		return http.StatusBadRequest
	case ErrMarriedStatusIDInvalid, ErrZipcodeRequired, ErrProvinceCodeInvalid, ErrBirthdayFormat:
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
	case ErrUsernameRequired:
		return "กรุณาระบุ อีเมล หรือ เบอร์โทรศัพท์"
	case ErrEmailInvalid:
		return "รูปแบบอีเมลไม่ถูกต้อง"
	case ErrPasswordRequires:
		return "กรุณากำหนดระหัสผ่าน"
	case ErrPasswordInvalid:
		return "พาสเวิร์ดต้องมีความยาว 6 - 20 ตัวอักษร"
	case ErrPasswordNotMatch:
		return "การยืนยันรหัสผ่านไม่ถูกต้อง"
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
	case ErrEmailDuplicated:
		return "ท่านไม่สามารถใช้อีเมลนี้ได้ เนื่องจากมีอีเมลดังกล่าวลงทะเบียนไว้แล้ว"
	case ErrPhoneDuplicated:
		return "ท่านไม่สามารถใช้เบอร์โทรศัพท์นี้ได้ เนื่องจากมีอเบอร์โทรศัพท์ดังกล่าวลงทะเบียนไว้แล้ว"
	case ErrSubDistrictRequired:
		return "ท่านไม่ได้ระบุข้อมูล ตำบล"
	case ErrDistrictRequired:
		return "ท่านไม่ได้ระบุข้อมูล อำเภอ"
	case ErrProvinceCodeRequired:
		return "ท่านไม่ได้ระบุข้อมูล จังหวัด"
	case ErrEmailRequired:
		return "กรุณาระบุ อีเมล์ หรือ เบอร์โทรศัพท์ เพื่อเข้าสู่ระบบ"
	case ErrInvalidCredentials:
		return "อีเมลหรือหมายเลขโทรศัพท์ที่คุณป้อนไม่ตรงกับบัญชีผู้ใช้ใดๆ"
	case ErrGenderIDInvalid:
		return "กรุณาระบุเพศของท่าน"
	case ErrMarriedStatusIDInvalid:
		return "กรุณาระบุสถานะการสมรสของท่าน"
	case ErrZipcodeRequired:
		return "กรุณาระบุรหัสไปรษณีย์สำหรับที่อยู่ของท่าน"
	case ErrProvinceCodeInvalid:
		return "ท่านระบุข้อมูลจังหวัดไม่ถูกต้อง"
	case ErrBirthdayFormat:
		return "format birthday not available [dd/mm/yyyy]"

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
