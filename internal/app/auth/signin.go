package auth

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/password"
	"strings"
	"time"
)

type signInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type signInResponse struct {
	Token  string `json:"token"`
	Name   string `json:"name"`
	RoleID int    `json:"roleID"`
}

func signIn(ctx context.Context, req signInRequest) (res signInResponse, err error) {
	req.Username = strings.TrimSpace(req.Username)
	req.Username = strings.ToLower(req.Username)

	if req.Username == "" {
		return res, ErrUsernameRequired
	}

	if req.Password == "" {
		return res, ErrPasswordRequired
	}

	var (
		userID         int64
		hashedPassword string
		name           string
		roleID         int
	)
	userID, hashedPassword, name, roleID, err = loginByEmail(ctx, req.Username)

	if err != nil {
		fmt.Println(err)
		userID, hashedPassword, name, roleID, err = loginByPhone(ctx, req.Username)
	}

	if err == sql.ErrNoRows {
		return res, ErrInvalidCredentials
	}

	if !password.Compare(hashedPassword, req.Password) {
		return res, ErrLoginFailed
	}

	token, err := generateToken()
	if err != nil {
		return res, err
	}

	/*
		map token กับ userID
	*/
	err = RedisClient.Set("creditScoring"+token, userID, 1*24*time.Hour).Err()
	if err != nil {
		return res, err
	}

	res.Token = token
	res.Name = name
	res.RoleID = roleID
	return
}

func loginByEmail(ctx context.Context, email string) (userID int64, hashedPassword string, name string, roleID int, err error) {
	err = dbctx.QueryRow(ctx, `
			select id, password, name, roleID
			from users
			where email = $1
	`, email).Scan(&userID, &hashedPassword, &name, &roleID)
	return
}

func loginByPhone(ctx context.Context, phone string) (userID int64, hashedPassword string, name string, roleID int, err error) {
	err = dbctx.QueryRow(ctx, `
			select id, password, name, roleID
			from users
			where phone = $1
	`, phone).Scan(&userID, &hashedPassword, &name, &roleID)
	return
}

func generateToken() (string, error) {
	var b [32]byte
	_, err := rand.Read(b[:])
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b[:]), nil
}
