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
	mux.HandleFunc("/update_home_address", updateHomeAddressHandler)
	mux.HandleFunc("/update_office_address", updateOfficeAddressHandler)
	mux.HandleFunc("/change_password", changePasswordHandler)
	mux.HandleFunc("/get_profile", getProfileHandler)
	mux.HandleFunc("/get_profile_by_id", getProfileByIDHandler)
	mux.HandleFunc("/change_forgot_password_question", changeForgotPasswordQuestionHandler)
	mux.HandleFunc("/get_old_forgot_password_and_answer", getOldForgotPasswordAndAnswerHandler)

	return mux

}

func updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req UserStruct
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := updateProfile(ctx, req)
	t.EncodeResult(w, res, err)
}

func updateHomeAddressHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req UserStruct
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := updateHomeAddress(ctx, req)
	t.EncodeResult(w, res, err)
}

func updateOfficeAddressHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req UserStruct
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := updateOfficeAddress(ctx, req)
	t.EncodeResult(w, res, err)
}

func changePasswordHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req changePasswordRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := changePassword(ctx, req)
	t.EncodeResult(w, res, err)
}

func getProfileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getProfileRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := getProfile(ctx, req)
	t.EncodeResult(w, res, err)
}

func getProfileByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getProfileRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := getProfileByID(ctx, req)
	t.EncodeResult(w, res, err)
}

func changeForgotPasswordQuestionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req changeForgotPasswordQuestionRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := changeForgotPasswordQuestion(ctx, req)
	t.EncodeResult(w, res, err)
}

func getOldForgotPasswordAndAnswerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getOldForgotPasswordAndAnswerRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := getOldForgotPasswordAndAnswer(ctx, req)
	t.EncodeResult(w, res, err)
}
