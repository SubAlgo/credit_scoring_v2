package user

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

type userStruct struct {
	UserID          int64  `json:"userID"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	GenderStatus    int    `json:"genderStatus"`
	Birthday        string `json:"birthday"`
	MarriedStatusID int    `json:"marriedStatusID"`
	Phone           string `json:"phone"`
	Religion        string `json:"religion"`
	Facebook        string `json:"facebook"`
	IG              string `json:"ig"`
	Line            string `json:"line"`
}

// update user profile
func (u *userStruct) updateProfile(ctx context.Context) (err error) {
	_, err = dbctx.Exec(ctx, `
		update users
		set name = $2,
			surname = $3,
			birthday = $4,
			marriedID = $5,
			facebook = $6,
			ig = $7,
			line = $8,
			genderID = $9,
			religion = $10
		where id = $1
	`, u.UserID, u.Name, u.Surname, u.Birthday, u.MarriedStatusID, u.Facebook, u.IG, u.Line, u.GenderStatus, u.Religion)
	return
}
