package domain

import (
	"time"
)

const (
	UserTable     = "User"
	UserInfoTable = "UserInfo"
)

type Users []*User
type User struct {
	ID           *uint       `gorm:"column:Id;primaryKey"`
	UserName     string      `gorm:"column:UserName;"`
	HashPassword string      `gorm:"column:Password;"`
	UserInfo     []*UserInfo `gorm:"foreignKey:UserId;references:Id"`
}

func (User) TableName() string {
	return UserTable
}

func (u *User) GetName() string {
	if u != nil {
		return u.UserName
	}
	return ""
}

func (u *User) CheckNil() *User {
	if u.ID != nil {
		return u
	}
	return nil
}

func (u *User) GetRoles() []string {
	roles := []string{}
	if u != nil {
		for _, v := range u.UserInfo {
			roles = append(roles, v.GetRole())
		}
	}
	return roles
}

type UserInfo struct {
	ID        *uint  `gorm:"column:Id;primaryKey"`
	UserID    uint   `gorm:"column:UserId;"`
	Role      string `gorm:"column:Role;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserInfo) TableName() string {
	return UserInfoTable
}

func (ui *UserInfo) GetRole() string {
	if ui != nil {
		return ui.Role
	}
	return ""
}
