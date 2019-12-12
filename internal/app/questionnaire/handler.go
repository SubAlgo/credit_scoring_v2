package questionnaire

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
	mux.HandleFunc("/answer", questionnaireAnswerHandler)
	mux.HandleFunc("/update", questionnaireLoanerUpdateHandler)
	mux.HandleFunc("/send", questionnaireLoanerSendHandler)

	mux.HandleFunc("/verify", questionnaireWorkerVerifyHandler)
	mux.HandleFunc("/worker_send", questionnaireWorkerSendHandler)
	mux.HandleFunc("/approve", questionnaireWorkerApproveHandler)

	mux.HandleFunc("/get_questionnaire_data", questionnaireGetDataHandler)

	return mux
}

func questionnaireAnswerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req QuestionnaireStruct
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := questionnaireAnswer(ctx, &req)
	t.EncodeResult(w, res, err)
}

func questionnaireLoanerUpdateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req QuestionnaireStruct
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := questionnaireLoanerUpdate(ctx, &req)
	t.EncodeResult(w, res, err)
}

func questionnaireLoanerSendHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req QuestionnaireStruct
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := questionnaireLoanerSend(ctx, &req)
	t.EncodeResult(w, res, err)
}

func questionnaireWorkerVerifyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req QuestionnaireStruct
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := questionnaireWorkerVerify(ctx, &req)
	t.EncodeResult(w, res, err)
}

func questionnaireWorkerSendHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req QuestionnaireStruct
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := questionnaireWorkerSend(ctx, &req)
	t.EncodeResult(w, res, err)
}

func questionnaireGetDataHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getDataArgs
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := questionnaireGetData(ctx, req)
	t.EncodeResult(w, res, err)
}

func questionnaireWorkerApproveHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req approveArgs
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := questionnaireWorkerApprove(ctx, &req)
	t.EncodeResult(w, res, err)
}
