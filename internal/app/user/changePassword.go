package user

import (
	"context"
	"database/sql"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/password"
	userPKG "github.com/subalgo/credit_scoring_v2/internal/pkg/user"
	"strings"
	"unicode/utf8"
)

type changePasswordRequest struct {
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

func changePassword(ctx context.Context, req changePasswordRequest) (res processResponse, err error) {
	userID := auth.GetUserID(ctx)

	if userID == 0 {
		return res, ErrUserNotLogin
	}

	req.NewPassword = strings.TrimSpace(req.NewPassword)
	if req.NewPassword == "" {
		return res, ErrPasswordRequired
	}

	if n := utf8.RuneCountInString(req.NewPassword); n < 6 || n > 20 {
		return res, ErrPasswordInvalid
	}

	if req.NewPassword != req.ConfirmPassword {
		return res, ErrConfirmPasswordNotMatch
	}

	hashedPassword, err := userPKG.GetHashPassword(ctx, userID)

	if err == sql.ErrNoRows {
		return res, ErrGetHashedPassword
	}

	if !password.Compare(hashedPassword, req.OldPassword) {
		return res, ErrOldPasswordInvalid
	}
	newHashPassword, err := password.Hash(req.NewPassword)
	if err != nil {
		return res, ErrHashingPassword
	}

	err = setPassword(ctx, userID, newHashPassword)

	if err != nil {
		return res, ErrChangePassword
	}

	res.Message = "เปลี่ยนรหัสผ่านสำเร็จ"
	return
}
