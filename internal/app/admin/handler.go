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

	mux.HandleFunc("/get_list_admin", getListAdminHandler)
	mux.HandleFunc("/get_list_employee", getListEmployeeHandler)
	mux.HandleFunc("/get_list_all_worker", getListAllWorkerHandler)
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

//get list
func getListAdminHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getUserListRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := getListAdmin(ctx, req)
	t.EncodeResult(w, res, err)
}

func getListEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getUserListRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := getListEmployee(ctx, req)
	t.EncodeResult(w, res, err)
}

func getListAllWorkerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getUserListRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := getListAllWorker(ctx, req)
	t.EncodeResult(w, res, err)
}
