package questionnaire

import (
	"context"
	"fmt"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
	"strings"
)

type loanerData struct {
	No            int    `json:"no"`
	LoanerID      int64  `json:"loanerID"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Email         string `json:"email"`
	SendAt        string `json:"sendAt"`
	UpdatedAt     string `json:"updatedAt"`
	UpdatedByID   int64  `json:"updatedByID"`
	UpdatedByName string `json:"updatedByName"`
}

type loanerList struct {
	Total int           `json:"total"`
	List  []*loanerData `json:"list"`
}

type getLoanerListRequest struct {
	statusID int
	name     string
	surname  string
}

func getLoanerList(ctx context.Context, req getLoanerListRequest) (res loanerList, err error) {
	rows, err := dbctx.Query(ctx, `
			SELECT 	users.id, users.name, users.surname, 
					to_char(q.sendAt, 'DD Mon YYYY เวลา HH:MM:SS') as sendAT,
					to_char(q.updatedAt, 'DD Mon YYYY เวลา HH:MM:SS') as updatedAT,
					q.updatedBy as updatedByID,
					(select users.name as updated_by from users where users.id = q.updatedBy) as updatedByName
			FROM users
			LEFT JOIN questionnaire AS q ON users.id = q.loanerID
			WHERE 
				q.statusID = $1 
				AND 
				users.name LIKE '%' || $2 || '%'
				AND
				users.surname LIKE '%' || $3 || '%'
			ORDER BY updatedAt ASC
		`, req.statusID, req.name, req.surname)

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	err = dbctx.QueryRow(ctx, `
			SELECT 	count(users.id)
			FROM users
			LEFT JOIN questionnaire AS q ON users.id = q.loanerID
			WHERE 
				q.statusID = $1 
				AND 
				users.name LIKE '%' || $2 || '%'
				AND
				users.surname LIKE '%' || $3 || '%'
		`, req.statusID, req.name, req.surname).Scan(&res.Total)

	if err != nil {
		return res, err
	}
	res.List = make([]*loanerData, 0)
	i := 1

	defer rows.Close()
	type handleNull struct {
		updateByName NullStringFromUpdateBy
	}
	for rows.Next() {
		var x loanerData
		var hd handleNull
		err = rows.Scan(
			&x.LoanerID, &x.Name, &x.Surname, &x.SendAt, &x.UpdatedAt, &x.UpdatedByID, &hd.updateByName,
		)

		if err != nil {
			return res, err
		}
		x.UpdatedByName = hd.updateByName.String

		x.No = i
		i = i + 1

		x.SendAt = setDateThaiFormat(x.SendAt)
		x.UpdatedAt = setDateThaiFormat(x.UpdatedAt)

		res.List = append(res.List, &x)
	}

	if err := rows.Err(); err != nil {
		return res, err
	}
	return
}

func getLoanerListDesc(ctx context.Context, req getLoanerListRequest) (res loanerList, err error) {
	rows, err := dbctx.Query(ctx, `
			SELECT 	users.id, users.name, users.surname, 
					to_char(q.sendAt, 'DD Mon YYYY เวลา HH:MM:SS') as sendAT,
					to_char(q.updatedAt, 'DD Mon YYYY เวลา HH:MM:SS') as updatedAT,
					q.updatedBy as updatedByID,
					(select users.name as updated_by from users where users.id = q.updatedBy) as updatedByName
			FROM users
			LEFT JOIN questionnaire AS q ON users.id = q.loanerID
			WHERE 
				q.statusID = $1 
				AND 
				users.name LIKE '%' || $2 || '%'
				AND
				users.surname LIKE '%' || $3 || '%'
			ORDER BY updatedAt DESC
		`, req.statusID, req.name, req.surname)

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	err = dbctx.QueryRow(ctx, `
			SELECT 	count(users.id)
			FROM users
			LEFT JOIN questionnaire AS q ON users.id = q.loanerID
			WHERE 
				q.statusID = $1 
				AND 
				users.name LIKE '%' || $2 || '%'
				AND
				users.surname LIKE '%' || $3 || '%'
		`, req.statusID, req.name, req.surname).Scan(&res.Total)

	if err != nil {
		return res, err
	}
	res.List = make([]*loanerData, 0)
	i := 1

	type loanerDataHandleNull struct {
		updateByName NullStringFromUpdateBy
	}

	defer rows.Close()
	for rows.Next() {
		var x loanerData
		var hd loanerDataHandleNull

		err = rows.Scan(
			&x.LoanerID, &x.Name, &x.Surname, &x.SendAt, &x.UpdatedAt, &x.UpdatedByID, &hd.updateByName,
		)
		if err != nil {
			fmt.Println(err)
			return res, err
		}
		x.SendAt = setDateThaiFormat(x.SendAt)
		x.UpdatedAt = setDateThaiFormat(x.UpdatedAt)
		x.UpdatedByName = hd.updateByName.String

		x.No = i
		i = i + 1

		res.List = append(res.List, &x)
	}

	if err := rows.Err(); err != nil {
		return res, err
	}
	return
}

func setDateThaiFormat(s string) string {
	var myStr []string
	myStr = strings.Split(s, " ")
	switch myStr[1] {
	case "Jan":
		s = myStr[0] + " มกราคม " + myStr[2] + " " + myStr[3] + " " + myStr[4]
	case "Feb":
		s = myStr[0] + " กุมภาพันธ์ " + myStr[2] + " " + myStr[3] + " " + myStr[4]
	case "Mar":
		s = myStr[0] + " มีนาคม " + myStr[2] + " " + myStr[3] + " " + myStr[4]
	case "Apr":
		s = myStr[0] + " เมษายน " + myStr[2] + " " + myStr[3] + " " + myStr[4]
	case "May":
		s = myStr[0] + " พฤษภาคม " + myStr[2] + " " + myStr[3] + " " + myStr[4]
	case "Jun":
		s = myStr[0] + " มิถุนายน " + myStr[2] + " " + myStr[3] + " " + myStr[4]
	case "Jul":
		s = myStr[0] + " กรกฏาคม " + myStr[2] + " " + myStr[3] + " " + myStr[4]
	case "Aug":
		s = myStr[0] + " สิงหาคม " + myStr[2] + " " + myStr[3] + " " + myStr[4]
	case "Sep", "Sept":
		s = myStr[0] + " กันยายน " + myStr[2] + " " + myStr[3] + " " + myStr[4]
	case "Oct":
		s = myStr[0] + " ตุลาคม " + myStr[2] + " " + myStr[3] + " " + myStr[4]
	case "Nov":
		s = myStr[0] + " พฤศจิกายน " + myStr[2] + " " + myStr[3] + " " + myStr[4]
	case "Dec":
		s = myStr[0] + " ธันวาคม " + myStr[2] + " " + myStr[3] + " " + myStr[4]
	default:
	}
	return s
}
