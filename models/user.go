package models

import (
	"time"
)

type User struct {
	ID          *uint     `json:"id"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	UserName    string    `json:"user_name"`
	Password    string    `json:"password,omitempty"`
	CreateDatim time.Time `json:"create_datim"`
}

type GetUserReq struct {
	UserEmail string `json:"user_email"`
}

func (r *GetUserReq) GetUserEmail() string {
	if r != nil {
		return r.UserEmail
	}
	return ""
}

type SignInResp struct {
	Token    string   `json:"token"`
	Role     []string `json:"role"`
	UserName string   `json:"user_name"`
}
