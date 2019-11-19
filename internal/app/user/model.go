package user

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type processResponse struct {
	Message string `json:"message"`
}

type UserStruct struct {
	UserID            int64  `json:"userID"`
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	GenderStatus      int    `json:"genderStatus"`
	Birthday          string `json:"birthday"`
	MarriedStatusID   int    `json:"marriedStatusID"`
	Phone             string `json:"phone"`
	Religion          string `json:"religion"`
	Facebook          string `json:"facebook"`
	IG                string `json:"ig"`
	Line              string `json:"line"`
	Address1          string `json:"address1"`
	Address2          string `json:"address2"`
	SubDistrict       string `json:"subDistrict"`
	District          string `json:"district"`
	ProvinceCode      string `json:"provinceCode"`
	ZipCode           string `json:"zipCode"`
	OfficeName        string `json:"officeName"`
	Address1Office    string `json:"address1Office"`
	Address2Office    string `json:"address2Office"`
	SubDistrictOffice string `json:"subDistrictOffice"`
	DistrictOffice    string `json:"districtOffice"`
	ProvinceOffice    string `json:"provinceOffice"`
	ZipCodeOffice     string `json:"zipCodeOffice"`
}

// update user profile
func (u *UserStruct) updateProfile(ctx context.Context) (err error) {
	_, err = dbctx.Exec(ctx, `
		update users
		set name = $2,
			surname = $3,
			birthday = $4,
			marriedID = $5,
			facebook = $6,
			ig = $7,
			line = $8,
			genderID = $9,
			religion = $10
		where id = $1
	`, u.UserID, u.Name, u.Surname, u.Birthday, u.MarriedStatusID, u.Facebook, u.IG, u.Line, u.GenderStatus, u.Religion)
	return
}

// update house address
func (u *UserStruct) updateHomeAddress(ctx context.Context) (err error) {
	_, err = dbctx.Exec(ctx, `
		update users
		set address1_home = $2,
			address2_home = $3,
			subDistrict_home = $4,
			district_home = $5,
			provinceCode_home = $6,
			zipCode_home = $7
		where id = $1
	`, u.UserID, u.Address1, u.Address2, u.SubDistrict, u.District, u.ProvinceCode, u.ZipCode)
	return
}
