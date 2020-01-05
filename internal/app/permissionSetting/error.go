package permissionSetting

import (
	"errors"
	"net/http"
)

var (
	ErrSignInRequired          = errors.New("sign in required")
	ErrPermissionNotAllow      = errors.New("permission not allow")
	ErrUpdatePermission        = errors.New("update permission error")
	ErrAccessPermissionGetData = errors.New("get permission error")
)

func errorToStatusCode(err error) int {
	switch err {
	case ErrSignInRequired, ErrPermissionNotAllow, ErrUpdatePermission:
		return http.StatusBadRequest
	case ErrAccessPermissionGetData:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}

func errorToMessage(err error) string {
	switch err {
	case ErrSignInRequired:
		return "ท่านยังไม่ได้เข้าสู่ระบบ"
	case ErrPermissionNotAllow:
		return "ท่านไม่มีสิทธิใช้งานฟังก์ชั่นนี้"
	case ErrUpdatePermission:
		return "เกิดการผิดพลาดในการอัพเดทข้อมูล"
	case ErrAccessPermissionGetData:
		return "เกิดข้อผิดพลาดในการเรียกข้อมูล permission"
	default:
		return "Internal Server Error"
	}
}
