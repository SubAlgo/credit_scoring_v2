package user

import (
	"context"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"strings"
)

func updateProfile(ctx context.Context, req UserStruct) (res processResponse, err error) {
	req.UserID = auth.GetUserID(ctx)

	if req.UserID == 0 {
		return res, ErrUserNotLogin
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Surname = strings.TrimSpace(req.Surname)
	req.Birthday = strings.TrimSpace(req.Birthday)
	req.Phone = strings.TrimSpace(req.Phone)
	req.Facebook = strings.TrimSpace(req.Facebook)
	req.IG = strings.TrimSpace(req.IG)
	req.Line = strings.TrimSpace(req.Line)

	if req.Name == "" {
		return res, ErrNameRequired
	}
	if req.Surname == "" {
		return res, ErrSurnameRequired
	}
	if req.Birthday == "" {
		return res, ErrBirthRequired
	}

	err = req.updateProfile(ctx)

	if err != nil {
		fmt.Println("error: ", err)
		return res, ErrUpdateProfile
	}

	res.Message = "บันทึกการเปลี่ยนแปลงข้อมูลส่วนตัวสำเร็จ"
	return
}
