package location

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type getLocationOPRequest struct {
	Zipcode string `json:"zipcode"`
}

type subDistrict struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type district struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type province struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type getLocationOPResponse struct {
	SubDistrictList []*subDistrict `json:"subDistrictList"`
	DistrictList    []*district    `json:"districtList"`
	ProvinceList    []*province    `json:"provinceList"`
}

func getLocationOptionByZipcode(ctx context.Context, req getLocationOPRequest) (res getLocationOPResponse, err error) {

	// Create SubDistrictList
	{
		subDistrictRows, err := dbctx.Query(ctx, `
								select subDistrict, subDistrict from location where zipCode = $1 group by subDistrict
							`, req.Zipcode)

		defer subDistrictRows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.SubDistrictList = make([]*subDistrict, 0)

		for subDistrictRows.Next() {
			var x subDistrict
			err = subDistrictRows.Scan(&x.Text, &x.Value)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.SubDistrictList = append(res.SubDistrictList, &x)
		}
		subDistrictRows.Close()
	}

	// Create DistrictList
	{
		districtRows, err := dbctx.Query(ctx, `
								select district, district from location where zipCode = $1 group by district
							`, req.Zipcode)

		defer districtRows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.DistrictList = make([]*district, 0)

		for districtRows.Next() {
			var x district
			err = districtRows.Scan(&x.Text, &x.Value)
			if err != nil {
				return res, err //ErrMakeSubDistrictList
			}
			res.DistrictList = append(res.DistrictList, &x)
		}
		districtRows.Close()
	}

	// Create ProvinceList
	{
		provinceRows, err := dbctx.Query(ctx, `
								select l.province_code, p.title
								from location as l 
								left join provinces as p on l.province_code = p.code
								where l.zipCode = $1 limit 1
							`, req.Zipcode)

		defer provinceRows.Close()

		if err != nil {
			return res, err //ErrGetSubDistrictRows
		}

		res.ProvinceList = make([]*province, 0)

		for provinceRows.Next() {
			var x province
			err = provinceRows.Scan(&x.Value, &x.Text)
			if err != nil {
				return res, ErrMakeSubDistrictList
			}
			res.ProvinceList = append(res.ProvinceList, &x)
		}
		provinceRows.Close()
	}

	return
}
