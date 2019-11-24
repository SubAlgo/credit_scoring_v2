package user

import (
	"database/sql"
	"reflect"
)

type processResponse struct {
	Message string `json:"message"`
}

type UserStruct struct {
	UserID             int64  `json:"userID"`
	Email              string `json:"email"`
	Name               string `json:"name"`
	Surname            string `json:"surname"`
	GenderID           int    `json:"genderID"`
	GenderStatus       string `json:"genderStatus"`
	Birthday           string `json:"birthday"`
	MarriedStatusID    int    `json:"marriedStatusID"`
	MarriedStatus      string `json:"marriedStatus"`
	Phone              string `json:"phone"`
	Religion           string `json:"religion"`
	Facebook           string `json:"facebook"`
	IG                 string `json:"ig"`
	Line               string `json:"line"`
	Address1           string `json:"address1"`
	Address2           string `json:"address2"`
	SubDistrict        string `json:"subDistrict"`
	District           string `json:"district"`
	ProvinceCode       string `json:"provinceCode"`
	ZipCode            string `json:"zipCode"`
	OfficeName         string `json:"officeName"`
	Address1Office     string `json:"address1Office"`
	Address2Office     string `json:"address2Office"`
	SubDistrictOffice  string `json:"subDistrictOffice"`
	DistrictOffice     string `json:"districtOffice"`
	ProvinceCodeOffice string `json:"provinceCodeOffice"`
	ZipCodeOffice      string `json:"zipCodeOffice"`
}

type userProfile struct {
	email             NullString
	name              NullString
	surname           NullString
	genderID          NullInt
	genderStatus      NullString
	marriedID         NullInt
	marriedStatus     NullString
	religion          NullString
	phone             NullString
	birthday          NullString
	child             NullInt
	facebook          NullString
	ig                NullString
	line              NullString
	address1Home      NullString
	address2Home      NullString
	subDistrictHome   NullString
	districtHome      NullString
	provinceHome      NullString
	zipCodeHome       NullString
	officeName        NullString
	address1Office    NullString
	address2Office    NullString
	subDistrictOffice NullString
	districtOffice    NullString
	provinceOffice    NullString
	zipCodeOffice     NullString
}

type NullString sql.NullString

func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then change String to "-" and Valid "true"
	if reflect.TypeOf(value) == nil {
		*ns = NullString{"-", true}
	} else {
		*ns = NullString{s.String, true}
	}
	return nil
}

type NullInt sql.NullInt64

func (ni *NullInt) Scan(value interface{}) error {
	var s sql.NullInt64
	if err := s.Scan(value); err != nil {
		return err
	}
	if reflect.TypeOf(value) == nil {
		*ni = NullInt{0, true}
	} else {
		*ni = NullInt{s.Int64, true}
	}
	return nil
}
