package admin

import (
	"context"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
)

func (c *createNewWorkerRequest) insert(ctx context.Context, hashPassword string) (id int64, err error) {
	err = dbctx.QueryRow(ctx, `
			insert into users
				(citizenID, email, password, name, surname, 
				birthday, phone, genderId, marriedId, religion, 
				roleId)
			values
				($1, $2, $3, $4, $5, 
				$6, $7, $8, $9, $10, 
				$11)
			returning id
		`, c.CitizenID, c.Email, hashPassword, c.Name, c.Surname,
		c.Birthday, c.Phone, c.GenderID, c.MarriedStatusID, c.Religion,
		c.RoleID).Scan(&id)
	if err != nil {
		return
	}
	return
}
