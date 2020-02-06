package admin

import (
	"errors"
	"net/http"
)

var (
	ErrEmailRequired      = errors.New("admin: email required")
	ErrPasswordRequires   = errors.New("admin: password required")
	ErrInvalidCredentials = errors.New("admin: invalid credentials")
	ErrEmailInvalid       = errors.New("admin: email invalid")
	ErrPasswordInvalid    = errors.New("admin: password invalid")
	ErrPasswordRequired   = errors.New("admin: password required")
	ErrHashPassword       = errors.New("admin: hash password has problem")
	ErrNameRequired       = errors.New("admin: name required")
	ErrSurNameRequired    = errors.New("admin: surname required")
	ErrPhoneLength        = errors.New("admin: phone number length not equal 10")
	ErrPhoneMustBeInt     = errors.New("admin: phone number must be integer")
	ErrEmailNotAvailable  = errors.New("admin: email not available")
	ErrPhoneNotAvailable  = errors.New("admin: phone not available")
	ErrSomething          = errors.New("admin: error something")
	ErrTokenRequired      = errors.New("admin: token required")
	ErrNotSignIn          = errors.New("admin: not sign in")

	ErrUsernameRequired   = errors.New("admin: required email or phone number")
	ErrPermissionNotAllow = errors.New("admin: permission not allow")
	ErrParamMissing       = errors.New("admin: param missing")
	ErrUpdateUserRole     = errors.New("admin: update user role")

	ErrGetListEmployee = errors.New("admin: get list employee")

	ErrGetWorkerRole          = errors.New("admin: error get worker")
	ErrDisableDeleteSuperUser = errors.New("admin: can not delete super user")
	ErrDeleteWorker           = errors.New("admin: error delete worker from DB")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrEmailRequired, ErrEmailInvalid, ErrPasswordRequires, ErrPasswordInvalid, ErrNameRequired, ErrSurNameRequired, ErrPhoneLength, ErrPhoneMustBeInt, ErrEmailNotAvailable, ErrPhoneNotAvailable:
		return http.StatusBadRequest

	case ErrUsernameRequired, ErrTokenRequired, ErrParamMissing:
		return http.StatusBadRequest
	case ErrInvalidCredentials, ErrNotSignIn:
		return http.StatusUnauthorized
	case ErrHashPassword, ErrSomething, ErrUpdateUserRole, ErrGetListEmployee, ErrGetWorkerRole:
		return http.StatusInternalServerError
	case ErrPermissionNotAllow:
		return http.StatusServiceUnavailable
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

	case ErrPermissionNotAllow:
		return "ท่านไม่มีสิทธิ์ใช้งานฟังค์ชั่นนี้"
	case ErrUpdateUserRole:
		return "เกิดปัญหาในการกำหนด role"
	case ErrParamMissing:
		return "ท่านไม่ได้ระบุรหัสพนักงาน หรือ รหัสสิทธิผู้ใช้งาน"

	case ErrSomething:
		return "Internal server error"
	case ErrGetListEmployee:
		return "Internal server error (getListEmployee)"
	case ErrGetWorkerRole:
		return "เกิดข้อผิดพลาดในการระบุข้อมูลพนักงานที่ต้องการลบ"
	case ErrDisableDeleteSuperUser:
		return "ไม่สามารถลบข้อมูลพนักงานระดับ super user ได้"
	case ErrDeleteWorker:
		return "เกิดข้อผิดพลาดในการลบข้อมูลพนักงาน"
	default:
		return "Internal Server Error"
	}

}
