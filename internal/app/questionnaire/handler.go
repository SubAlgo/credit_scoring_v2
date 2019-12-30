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
	mux.HandleFunc("/check_status", questionnaireCheckStatusHandler)
	mux.HandleFunc("/check_status_by_id", questionnaireCheckStatusByIDHandler)

	mux.HandleFunc("/answer", questionnaireAnswerHandler)
	mux.HandleFunc("/update", questionnaireLoanerUpdateHandler)
	mux.HandleFunc("/send", questionnaireLoanerSendHandler)

	mux.HandleFunc("/verify", questionnaireWorkerVerifyHandler)
	mux.HandleFunc("/worker_send", questionnaireWorkerSendHandler)
	mux.HandleFunc("/approve", questionnaireWorkerApproveHandler)
	mux.HandleFunc("/deny", questionnaireWorkerDenyHandler)

	mux.HandleFunc("/get_questionnaire_data", questionnaireGetDataHandler)

	mux.HandleFunc("/get_list_new_loaner", questionnaireGetListNewLoanerHandler)
	mux.HandleFunc("/get_list_in_verify", questionnaireGetListInVerifyHandler)
	mux.HandleFunc("/get_list_wait_approve", questionnaireGetListWaitApproveHandler)

	mux.HandleFunc("/get_approve_result", questionnaireGetApproveResultHandler)

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

func questionnaireWorkerDenyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req approveArgs
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := questionnaireWorkerDeny(ctx, &req)
	t.EncodeResult(w, res, err)
}

func questionnaireGetListNewLoanerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getQuestionnaireListRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := questionnaireGetListNewLoaner(ctx, req)
	t.EncodeResult(w, res, err)
}

func questionnaireGetListInVerifyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getQuestionnaireListRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := questionnaireGetListInVerify(ctx, req)
	t.EncodeResult(w, res, err)
}

func questionnaireGetListWaitApproveHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getQuestionnaireListRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := questionnaireGetListWaitApprove(ctx, req)
	t.EncodeResult(w, res, err)
}

func questionnaireCheckStatusHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req checkQuestionnaireStatusRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := questionnaireCheckStatus(ctx, req)
	t.EncodeResult(w, res, err)
}

func questionnaireCheckStatusByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req checkQuestionnaireStatusRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := questionnaireCheckStatusByID(ctx, req)
	t.EncodeResult(w, res, err)
}

func questionnaireGetApproveResultHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getApproveResultRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := questionnaireGetApproveResult(ctx, req)
	t.EncodeResult(w, res, err)
}
