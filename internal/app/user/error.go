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
	ErrAddress1Required          = errors.New("updateHomeAddress: address1 invalid")
	ErrAddress2Required          = errors.New("updateHomeAddress: address2 invalid")
	ErrSubDistrictRequired       = errors.New("updateHomeAddress: sub district invalid")
	ErrDistrictRequired          = errors.New("updateHomeAddress: district invalid")
	ErrProvinceCodeRequired      = errors.New("updateHomeAddress: province code invalid")
	ErrZipcodeRequired           = errors.New("updateHomeAddress: zipCode invalid")
	ErrUpdateDataBaseHomeAddress = errors.New("updateHomeAddress: update database error")
)

var (
	ErrProvinceCodeInvalid         = errors.New("updateOfficeAddress: province code invalid")
	ErrZipcodeOfficeRequired       = errors.New("updateOfficeAddress: zipcode invalid")
	ErrUpdateDataBaseOfficeAddress = errors.New("updateOfficeAddress: update database")
)

// changePassword
var (
	ErrConfirmPasswordNotMatch = errors.New("changePassword: new password not match with confirm password")
	ErrPasswordRequired        = errors.New("changePassword: password required")
	ErrPasswordInvalid         = errors.New("changePassword: password invalid")
	ErrGetHashedPassword       = errors.New("changePassword: get hashed password error")
	ErrOldPasswordInvalid      = errors.New("changePassword: old password invalid")
	ErrHashingPassword         = errors.New("changePassword: hashing password")
	ErrChangePassword          = errors.New("changePassword: update database")
)

// get profile
var (
	ErrGetProfile       = errors.New("getProfile: get profile error")
	ErrGetProfileNoRows = errors.New("getProfile: sql not row")
	ErrUserIdRequired   = errors.New("getProfileByID: userID required")
	ErrNotPermission    = errors.New("getProfileByID: permission not allow")
	ErrPkgGetAge        = errors.New("getProfile: package user.GetAge")
)

// set new forgot password
var (
	ErrQuestionIDRequired                    = errors.New("forgotPasswordQuestion: question id required")
	ErrForgotPasswordAnswer                  = errors.New("forgotPasswordQuestion: question answer required")
	ErrUpdateForgotPasswordQuestion          = errors.New("forgotPasswordQuestion: update to database")
	ErrSelectForgotPasswordQuestionAndAnswer = errors.New("forgotPasswordQuestion: select question and answer")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrUserNotLogin:
		return http.StatusNotAcceptable
	case ErrNameRequired, ErrSurnameRequired, ErrPhoneRequired, ErrBirthRequired:
		return http.StatusBadRequest
	case ErrAddress1Required, ErrAddress2Required, ErrSubDistrictRequired, ErrDistrictRequired, ErrProvinceCodeRequired, ErrZipcodeRequired:
		return http.StatusBadRequest
	case ErrZipcodeOfficeRequired, ErrConfirmPasswordNotMatch, ErrPasswordRequired, ErrPasswordInvalid, ErrOldPasswordInvalid:
		return http.StatusBadRequest
	case ErrUpdateProfile, ErrUpdateDataBaseHomeAddress, ErrUpdateDataBaseOfficeAddress, ErrGetHashedPassword, ErrHashingPassword, ErrChangePassword:
		return http.StatusInternalServerError
	case ErrGetProfile, ErrGetProfileNoRows, ErrPkgGetAge:
		return http.StatusInternalServerError
	case ErrUserIdRequired:
		return http.StatusBadRequest
	case ErrNotPermission:
		return http.StatusNotAcceptable
	case ErrQuestionIDRequired, ErrForgotPasswordAnswer:
		return http.StatusBadRequest
	case ErrUpdateForgotPasswordQuestion, ErrSelectForgotPasswordQuestionAndAnswer:
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
	case ErrZipcodeOfficeRequired:
		return "กรุณาระบุรหัสไปรษณีของสถานที่ทำงานของท่าน"
	case ErrProvinceCodeInvalid:
		return "ระบุรหัสจังหวัดไม่ถูกต้อง"
	case ErrConfirmPasswordNotMatch:
		return "ยืนยันรหัสผ่านที่ต้องการเปลี่ยนไม่ถูกต้อง"
	case ErrPasswordRequired:
		return "กรุณาระบุรหัสผ่านที่ต้องการเปลี่ยน"
	case ErrPasswordInvalid:
		return "รหัสผ่านต้องมีความยาวระหว่าง 6-20 ตัวอักษร"
	case ErrOldPasswordInvalid:
		return "ระบุรหัสผ่านเดิมไม่ถูกต้อง"
	case ErrUserIdRequired:
		return "user id required"
	case ErrGetProfileNoRows:
		return "ไม่พบข้อมูลของผู้ใช้รายนี้"
	case ErrNotPermission:
		return "ท่านไม่สามารถใช้ฟังค์ชั่นนี้ได้"
	case ErrQuestionIDRequired:
		return "กรุณาระบุคำถามสำหรับกู้รหัสผ่าน"
	case ErrForgotPasswordAnswer:
		return "กรุณาระบุคำตอบสำหรับกู้รหัสผ่าน"
	case ErrUpdateForgotPasswordQuestion:
		return "internal server error (change forgot password question)"
	case ErrSelectForgotPasswordQuestionAndAnswer:
		return "internal server error (select forgot password question and answer)"
	default:
		return "internal server error"
	}
}
