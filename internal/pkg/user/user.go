package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/acoshift/pgsql"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

var (
	ErrEmailDuplicated = errors.New("user: email duplicated")
	ErrNotFound        = errors.New("user: not found")
	ErrPhoneDuplicated = errors.New("user: phone number duplicated")
)

type SignUpArgs struct {
	Name            string
	Surname         string
	Email           string
	Phone           string
	HashPassword    string
	Birthday        string
	GenderID        int
	MarriedStatusID int
	Religion        string
	RoleId          int
}

func Insert(ctx context.Context, arg *SignUpArgs) (id int64, err error) {
	err = dbctx.QueryRow(ctx,
		`insert into users 
				(email, password, name, surname, birthday, phone, genderId , marriedId, religion, roleId)
			values
				($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			returning id
		`, arg.Email, arg.HashPassword, arg.Name, arg.Surname, arg.Birthday, arg.Phone, arg.GenderID, arg.MarriedStatusID, arg.MarriedStatusID, arg.RoleId).Scan(&id)

	if pgsql.IsUniqueViolation(err, "users_email_idx") {
		return 0, ErrEmailDuplicated
	}

	if pgsql.IsUniqueViolation(err, "users_phone_idx") {
		return 0, ErrPhoneDuplicated
	}

	if err != nil {
		return 0, err //ErrEmailDuplicated
	}
	return
}

func GetUserRole(ctx context.Context, userID int64) (roleID int, err error) {
	err = dbctx.QueryRow(ctx, `
			select roleID
			from users
			where id = $1
	`, userID).Scan(&roleID)
	if err == sql.ErrNoRows {
		return 0, ErrNotFound
	}
	return
}
