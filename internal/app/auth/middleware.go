package auth

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/user"
	"net/http"
	"time"
)

type (
	ctxKeyUserID struct{}   // store user id
	ctxKeyUserRole struct{} // store user role
	ctxKeyToken struct{}
)

func FetchAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// X-Token
		token := r.Header.Get("X-Token")

		//ถ้า token เป็นค่าว่างไป middleware ถัดไป (call next)
		if token == "" {
			h.ServeHTTP(w, r)
			return
		}

		key := "creditScoring" + token
		userID, err := RedisClient.Get(key).Int64()
		if err != nil {
			h.ServeHTTP(w, r)
			return
		}
		// set new redis expire
		err = RedisClient.Expire(key, 1*24*time.Hour).Err()
		if err != nil {
			h.ServeHTTP(w, r)
			return
		}
		ctx := r.Context()

		/*
			Set user role to context
			1. get user role id from DB
			2. set context value by key:ctxKeyUserRole and value: userRoleID
		*/
		userRoleID, err := user.GetUserRole(ctx, userID)
		if err != nil {
			h.ServeHTTP(w, r)
			return
		}

		/*
			สร้าง context สำหรับเก็บ userID, userRole
		*/
		ctx = context.WithValue(ctx, ctxKeyToken{}, token)
		ctx = context.WithValue(ctx, ctxKeyUserID{}, userID)
		ctx = context.WithValue(ctx, ctxKeyUserRole{}, userRoleID)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

//GetUserID get user id by context
func GetUserID(ctx context.Context) int64 {
	userID, _ := ctx.Value(ctxKeyUserID{}).(int64)
	return userID
}

//GetUserRole get user role by context
func GetUserRole(ctx context.Context) int {
	roleID, _ := ctx.Value(ctxKeyUserRole{}).(int)
	return roleID
}
