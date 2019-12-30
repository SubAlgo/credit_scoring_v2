package province

import (
	"github.com/subalgo/credit_scoring_v2/internal/pkg/transport"
	"net/http"
)

var t = transport.HTTP{
	ErrorToStatusCode: errorToStatusCode,
	ErrorToMessage:    errorToMessage,
}

func Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/set_province_score", setProvinceScoreHandler)
	mux.HandleFunc("/get_province_list_by_geo_id", getProvinceListHandler)
	
	return mux
}

func getProvinceListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getProvinceListRequest

	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := getProvinceListByGeoID(ctx, req)
	t.EncodeResult(w, res, err)
}

func setProvinceScoreHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req setProvinceScoreRequest

	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := setProvinceScore(ctx, req)
	t.EncodeResult(w, res, err)
}
