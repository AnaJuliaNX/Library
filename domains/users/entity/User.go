package entity

import (
	PKGJWT "library/PKGJWT"
)

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Owner       bool   `json:"owner"`
	Phone       string `json:"phone"`
	DateCreated string `json:"date_created"`
}

func (u User) GenerateToken() (string, int64, error) {
	return PKGJWT.GenerateToken(
		u.ID,
		u.Email,
		u.Owner,
		u.DateCreated,
	)
}

var UpdateUserValidate = map[string]string{
	"name":  "type:string",
	"email": "type:string",
	"phone": "type:string",
}
