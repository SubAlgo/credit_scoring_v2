package questionnaireOption

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
	mux.HandleFunc("/get_option", getOptionHandler)

	return mux
}

func getOptionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getQuestionnaireOptionRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := getOption(ctx, req)
	t.EncodeResult(w, res, err)
}
