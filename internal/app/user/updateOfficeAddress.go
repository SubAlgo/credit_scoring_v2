package user

import (
	"context"
	"fmt"
	"github.com/acoshift/pgsql"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"strings"
)

func updateOfficeAddress(ctx context.Context, req UserStruct) (res processResponse, err error) {
	req.UserID = auth.GetUserID(ctx)

	if req.UserID == 0 {
		return res, ErrUserNotLogin
	}

	req.OfficeName = strings.TrimSpace(req.OfficeName)
	req.Address1Office = strings.TrimSpace(req.Address1Office)
	req.Address2Office = strings.TrimSpace(req.Address2Office)
	req.SubDistrictOffice = strings.TrimSpace(req.SubDistrictOffice)
	req.DistrictOffice = strings.TrimSpace(req.DistrictOffice)
	req.ProvinceCodeOffice = strings.TrimSpace(req.ProvinceCodeOffice)

	if req.OfficeName == "" {
		req.OfficeName = "-"
	}

	if req.Address1Office == "" {
		req.Address1Office = "-"
	}

	if req.Address2Office == "" {
		req.Address2Office = "-"
	}

	if req.SubDistrictOffice == "" {
		req.SubDistrictOffice = "-"
	}

	if req.DistrictOffice == "" {
		req.DistrictOffice = "-"
	}

	if req.ProvinceCodeOffice == "" {
		req.ProvinceCodeOffice = "-"
	}

	if req.ZipCodeOffice == "" {
		//return res, ErrZipcodeOfficeRequired
	}

	err = req.updateOfficeAddress(ctx)

	if pgsql.IsUniqueViolation(err, "users_provincecode_office_fkey") {
		return res, ErrProvinceCodeInvalid
	}

	if err != nil {
		fmt.Println(err)
		return res, ErrUpdateDataBaseOfficeAddress
	}

	res.Message = "บันทึการเปลี่ยนแปลงข้อมูลสถานที่ทำงานสำเร็จ"

	return
}
