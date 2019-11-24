package admin

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
	mux.HandleFunc("/create_new_worker", createNewWorkerHandler)
	mux.HandleFunc("/change_user_role", changeUserRoleHandler)

	return mux
}

func createNewWorkerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req createNewWorkerRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := createNewWorker(ctx, req)
	t.EncodeResult(w, res, err)
}

func changeUserRoleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req changeUserRoleRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := changeUserRole(ctx, req)
	t.EncodeResult(w, res, err)
}
