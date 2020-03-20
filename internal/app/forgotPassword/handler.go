package forgotPassword

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
	mux.HandleFunc("/get_question_option", getForgotPasswordQuestionOptionHandler)
	mux.HandleFunc("/get_question", getQuestionByEmailHandler)
	mux.HandleFunc("/set_new_password", setNewPasswordHandler)
	return mux
}

func getForgotPasswordQuestionOptionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req forgotPasswordRequest

	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := getForgotPasswordQuestionOption(ctx, req)
	t.EncodeResult(w, res, err)
}

func getQuestionByEmailHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getQuestionByEmailRequest

	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := getQuestionByEmail(ctx, req)
	t.EncodeResult(w, res, err)
}

func setNewPasswordHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req setNewPasswordRequest

	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := setNewPassword(ctx, req)
	t.EncodeResult(w, res, err)
}
