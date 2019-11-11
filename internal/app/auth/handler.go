package auth

import (
	"github.com/go-redis/redis"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/transport"
	"net/http"
)

var t = transport.HTTP{
	ErrorToStatusCode: errorToStatusCode,
	ErrorToMessage:    errorToMessage,
}

var RedisClient *redis.Client

// Handler creates new http handler
func Handler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.NotFoundHandler())
	mux.HandleFunc("/signup", signUpHandler)
	mux.HandleFunc("/signin", signInHandler)
	mux.HandleFunc("/signout", signOutHandler)
	//mux.HandleFunc("/show", showHandler)
	return mux
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req signUpRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := signUp(ctx, req)
	t.EncodeResult(w, res, err)
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req signInRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := signIn(ctx, req)
	t.EncodeResult(w, res, err)
}

func signOutHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req signOutRequest
	err := t.DecodeRequest(w, r, &req)
	if err != nil {
		return
	}
	res, err := signOut(ctx, req)
	t.EncodeResult(w, res, err)
}
