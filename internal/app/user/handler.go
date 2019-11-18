package user

import (
	"github.com/subalgo/credit_scoring_v2/internal/pkg/transport"
	//"github.com/subalgo/credit_scoring_v2/internal/pkg/user"
	"net/http"
)

var t = transport.HTTP{
	ErrorToStatusCode: errorToStatusCode,
	ErrorToMessage:    errorToMessage,
}

func Handler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.NotFoundHandler())
	mux.HandleFunc("/update_profile", updateProfileHandler)

	return mux

}

func updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req userStruct
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := updateProfile(ctx, req)
	t.EncodeResult(w, res, err)
}
