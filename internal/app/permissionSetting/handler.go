package permissionSetting

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
	mux.HandleFunc("/update_access_permission", updateAccessPermissionHandler)
	mux.HandleFunc("/get_access_permission", getAccessPermissionHandler)

	return mux
}

func updateAccessPermissionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req accessPermissionArgs

	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := updateAccessPermission(ctx, req)
	t.EncodeResult(w, res, err)
}

func getAccessPermissionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getAccessPermissionRequest

	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := getAccessPermission(ctx, req)
	t.EncodeResult(w, res, err)
}
