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
	UserID uint `json:"user_id"`
}

func (r *GetUserReq) GetUserID() uint {
	if r != nil {
		return r.UserID
	}
	return 0
}

type SignInResp struct {
	Token    string   `json:"token"`
	Role     []string `json:"role"`
	UserName string   `json:"user_name"`
}
