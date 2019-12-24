package user

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
	userPK "github.com/subalgo/credit_scoring_v2/internal/pkg/user"
)

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

// update office address
func (u *UserStruct) updateOfficeAddress(ctx context.Context) (err error) {
	_, err = dbctx.Exec(ctx, `
		update users
		set office_name = $2,
			address1_office = $3,
			address2_office = $4,
			subDistrict_office = $5,
			district_office = $6,
			provinceCode_office = $7,
			zipCode_office = $8
		where id = $1
	`, u.UserID, u.OfficeName, u.Address1Office, u.Address2Office, u.SubDistrictOffice, u.DistrictOffice, u.ProvinceCodeOffice, u.ZipCodeOffice)
	return
}

func setPassword(ctx context.Context, userID int64, hashedPassword string) (err error) {
	_, err = dbctx.Exec(ctx, `
		update users
		set password = $2
		where id = $1
	`, userID, hashedPassword)
	return
}

// get user profile
func (u *UserStruct) getProfile(ctx context.Context, userID int64) (err error) {
	var up userProfile
	err = dbctx.QueryRow(ctx, `
		select 	email, name, surname, genderID, genderStatus.title, marriedID, marriedStatus.title, religion, birthday, phone, child,
				facebook, ig, line, 
				address1_home, address2_home, subDistrict_home, district_home, provinceCode_home, (select p.title from provinces as p where p.code = users.provinceCode_home), zipCode_home,
				office_name, address1_office, address2_office, subDistrict_office, district_office, provinceCode_office, (select p.title from provinces as p where p.code = users.provinceCode_office), zipCode_office
		from users
		left join marriedStatus on users.marriedID = marriedStatus.id
		left join genderStatus on users.genderID = genderStatus.id
		where users.id = $1
	`, userID).Scan(&up.email, &up.name, &up.surname, &up.genderID, &up.genderStatus, &up.marriedID, &up.marriedStatus, &up.religion, &up.birthday, &up.phone, &up.child,
		&up.facebook, &up.ig, &up.line,
		&up.address1Home, &up.address2Home, &up.subDistrictHome, &up.districtHome, &up.provinceHome, &up.provinceHomeTitle, &up.zipCodeHome,
		&up.officeName, &up.address1Office, &up.address2Office, &up.subDistrictOffice, &up.districtOffice, &up.provinceOffice, &up.provinceTitleOffice, &up.zipCodeOffice)

	if err != nil {
		return err
	}
	u.UserID = userID
	u.Email = up.email.String
	u.Name = up.name.String
	u.Surname = up.surname.String
	u.GenderID = int(up.genderID.Int64)
	u.GenderStatus = up.genderStatus.String
	u.Religion = up.religion.String
	u.Phone = up.phone.String
	u.Birthday = up.birthday.String
	u.MarriedStatusID = int(up.marriedID.Int64)
	u.MarriedStatus = up.marriedStatus.String
	u.Facebook = up.facebook.String
	u.IG = up.ig.String
	u.Line = up.line.String
	u.Address1 = up.address1Home.String
	u.Address2 = up.address2Home.String
	u.SubDistrict = up.subDistrictHome.String
	u.District = up.districtHome.String
	u.ProvinceCode = up.provinceHome.String
	u.ProvinceTitle = up.provinceHomeTitle.String
	u.ZipCode = up.zipCodeHome.String
	u.OfficeName = up.officeName.String
	u.Address1Office = up.address1Office.String
	u.Address2Office = up.address2Office.String
	u.SubDistrictOffice = up.subDistrictOffice.String
	u.DistrictOffice = up.districtOffice.String
	u.ProvinceCodeOffice = up.provinceOffice.String
	u.ProvinceTitleOffice = up.provinceTitleOffice.String
	u.ZipCodeOffice = up.zipCodeOffice.String
	u.Age, err = userPK.GetAge(ctx, userID)
	return
}
