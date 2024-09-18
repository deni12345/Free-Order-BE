package domain

import "time"

type Sheet struct {
	ID       *uint      `gorm:"column:Id;"`
	Name     string     `gorm:"column:SheetName;"`
	CreateAt *time.Time `gorm:"column:CreateAt;"`
	UserId   *uint      `gorm:"column:UserId;"`
	User     *User      `gorm:"foreignKet:UserId;references:Id"`
	OrderId  *uint      `gorm:"column:OrderId;"`
	Order    *Order     `gorm:"foreignKet:OrderId;references:Id"`
	IsActive bool       `gorm:"column:IsActive;"`
}
