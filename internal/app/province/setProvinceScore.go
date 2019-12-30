package province

import (
	"context"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type setProvinceScoreRequest struct {
	ProvinceCode  int `json:"provinceCode"`
	ProvinceScore int `json:"provinceScore"`
}

type setProvinceScoreResponse struct {
	Message string `json:"message"`
}

func setProvinceScore(ctx context.Context, req setProvinceScoreRequest) (res setProvinceScoreResponse, err error) {

	userID := auth.GetUserID(ctx)
	if userID == 0 {
		return res, ErrUserNotLogin
	}

	if req.ProvinceScore == 0 {
		return res, ErrInvalidProvinceScore
	}

	if req.ProvinceCode == 0 {
		return res, ErrInvalidProvinceID
	}

	{
		_, err = dbctx.Exec(ctx, `
		update provinces
		set score = $2
		where code = $1
		`, req.ProvinceCode, req.ProvinceScore)
	}
	if err != nil {
		fmt.Println("setProvinceScore: ", err)
		return res, err
	}

	res.Message = "บันทึกคะแนนจังหวัดสำเร็จ"
	return
}
