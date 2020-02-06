package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/acoshift/pgsql"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
	"time"
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

// insert new customer and create new worker
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

// delete user
func Delete(ctx context.Context, userID int64) (err error) {
	_, err = dbctx.Exec(ctx,
		`delete from users where id = $1
		`, userID)
	return
}

// get user role for store in context
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

// get hash password
func GetHashPassword(ctx context.Context, userID int64) (hashedPassword string, err error) {
	err = dbctx.QueryRow(ctx, `
			select password
			from users
			where id = $1
	`, userID).Scan(&hashedPassword)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAge(ctx context.Context, userID int64) (age int, err error) {
	var birthday string
	err = dbctx.QueryRow(ctx, `
		select birthday
		from users
		where id = $1
	`, userID).Scan(&birthday)

	layOut := "02/01/2006"
	dobTime, err := time.Parse(layOut, birthday) //แปลง format input to date type
	if err != nil {
		return 0, err
	}

	birthMonth := monthToInt(dobTime.Month().String())
	nowMonth := monthToInt(time.Now().Month().String())

	defMonth := nowMonth - birthMonth
	var minusAge int
	if defMonth < 0 {
		minusAge = -1
	} else {
		minusAge = 0
	}

	age = (time.Now().Year() - dobTime.Year()) + minusAge
	return
}
