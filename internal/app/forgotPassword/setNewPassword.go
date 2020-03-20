package forgotPassword

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/password"
	"math/rand"
	"strings"
	"time"
)

type setNewPasswordRequest struct {
	Email  string `json:"email"`
	Answer string `json:"answer"`
}

type setNewPasswordResponse struct {
	NewPassword string `json:"newPassword"`
}

/*
	function รับ email กับ forgotPasswordAnswer มาเทียบกัน ถ้าถูกจะ set new password and return newPassword
*/

// ใน db อาจจะเก็บคำถามกู้รหัสผ่านเป็น string เลย
func setNewPassword(ctx context.Context, req setNewPasswordRequest) (res setNewPasswordResponse, err error) {
	req.Email = strings.TrimSpace(req.Email)
	req.Answer = strings.TrimSpace(req.Answer)

	if req.Email == "" {
		return res, ErrEmailRequired
	}
	if req.Answer == "" {
		return res, ErrAnswerRequired
	}
	var id int64
	err = dbctx.QueryRow(ctx, `
				select id 
				from users
				where 
					email = $1 and forgotPassWordAns = $2
			`, req.Email, req.Answer).Scan(&id)
	_ = id

	if err == sql.ErrNoRows {
		return res, ErrEmailAndAnswerNotMarch
	}

	if err != nil {
		fmt.Println(err)
		return res, ErrEmailAndAnswerNotMarch
	}
	res.NewPassword = randomString(10)

	hashedPassword, _ := password.Hash(res.NewPassword)

	fmt.Println("newPassword: ", res.NewPassword)
	fmt.Println("hashedPassword: ", hashedPassword)

	_, err = dbctx.Exec(ctx, `
				update users
				set password = $1
				where
					email = $2
			`, hashedPassword, req.Email)

	if err != nil {
		fmt.Println(err)
		return res, ErrUpdateNewPassword
	}
	return
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func randomString(length int) string {
	return StringWithCharset(length, charset)
}
