package models

import (
	"time"
)

type User struct {
	ID          *string   `json:"id"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	UserName    string    `json:"user_name"`
	Password    string    `json:"password,omitempty"`
	IsActive    bool      `json:"is_active"`
	GoogleID    string    `json:"google_id,omitempty"`
	CreateDatim time.Time `json:"create_datim"`
}

type GoogelUser struct {
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
	Token    string `json:"token"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}
