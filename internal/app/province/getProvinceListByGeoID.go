package province

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type getProvinceListRequest struct {
	GeoID int `json:"geoID"`
}

type provinceData struct {
	ProvinceCode int    `json:"provinceCode"`
	Title        string `json:"title"`
	Score        int    `json:"score"`
}

type getProvinceListResponse struct {
	ProvinceList []*provinceData `json:"provinceList"`
}

func getProvinceListByGeoID(ctx context.Context, req getProvinceListRequest) (res getProvinceListResponse, err error) {
	userID := auth.GetUserID(ctx)

	if userID == 0 {
		return res, ErrUserNotLogin
	}
	if req.GeoID == 0 {
		req.GeoID = 1
	}
	if req.GeoID > 6 {
		req.GeoID = 6
	}

	rows, err := dbctx.Query(ctx, `
		select code, title, score
		from provinces
		where geographyID = $1
		order by title
	`, req.GeoID)

	if err != nil {
		return res, ErrSelectProvinceList
	}

	defer rows.Close()

	res.ProvinceList = make([]*provinceData, 0)

	for rows.Next() {
		var x provinceData
		err = rows.Scan(&x.ProvinceCode, &x.Title, &x.Score)

		if err != nil {
			return res, ErrCreateProvinceList
		}
		res.ProvinceList = append(res.ProvinceList, &x)
	}
	if err := rows.Err(); err != nil {
		return res, err
	}
	rows.Close()
	return
}
