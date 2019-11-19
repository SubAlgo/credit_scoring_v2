package user

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"strings"
)

func updateHomeAddress(ctx context.Context, req UserStruct) (res processResponse, err error) {
	req.UserID = auth.GetUserID(ctx)

	if req.UserID == 0 {
		return res, ErrUserNotLogin
	}

	req.Address1 = strings.TrimSpace(req.Address1)
	req.Address2 = strings.TrimSpace(req.Address2)
	req.SubDistrict = strings.TrimSpace(req.SubDistrict)
	req.District = strings.TrimSpace(req.District)
	req.ProvinceCode = strings.TrimSpace(req.ProvinceCode)

	if req.Address1 == "" {
		return res, ErrAddress1Required
	}

	if req.SubDistrict == "" {
		return res, ErrSubDistrictRequired
	}

	if req.District == "" {
		return res, ErrDistrictRequired
	}

	if req.ProvinceCode == "" {
		return res, ErrProvinceCodeRequired
	}

	if req.ZipCode == "" {
		return res, ErrZipcodeRequired
	}

	err = req.updateHomeAddress(ctx)

	if err != nil {
		return res, err
	}

	res.Message = "บันทึกข้อมูลการเปลี่ยนแปลงที่อยู่สำเร็จ"

	return
}
