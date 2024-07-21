package domain

import "time"

const (
	UserTable     = "User"
	UserInfoTable = "UserInfo"
)

type Users []*User
type User struct {
	ID       *uint    `gorm:"column:Id;primaryKey"`
	UserName string   `gorm:"column:UserName;"`
	Password string   `gorm:"column:Password;"`
	UserInfo UserInfo `gorm:"foreignKey:UserName;references:UserId"`
}

func (User) TableName() string {
	return UserTable
}

func (mu *User) CheckNil() *User {
	if mu.ID != nil {
		return mu
	}
	return nil
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
