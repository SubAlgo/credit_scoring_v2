package user

type UserStruct struct {
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	GenderStatus      int    `json:"genderStatus"`
	Birth             string `json:"birth"`
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
	Province          string `json:"province"`
	ZipCode           string `json:"zipCode"`
	OfficeName        string `json:"officeName"`
	Address1Office    string `json:"address1"`
	Address2Office    string `json:"address2"`
	SubDistrictOffice string `json:"subDistrict"`
	DistrictOffice    string `json:"district"`
	ProvinceOffice    string `json:"province"`
	ZipCodeOffice     string `json:"zipCode"`
}
