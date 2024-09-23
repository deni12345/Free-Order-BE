package domain

import "time"

const (
	SheetTable = "Sheet"
)

type Sheets []*Sheet

type Sheet struct {
	ID       *uint      `gorm:"column:Id;"`
	Name     string     `gorm:"column:SheetName;"`
	CreateAt *time.Time `gorm:"column:CreateAt;"`
	UserId   *uint      `gorm:"column:UserId;"`
	User     *User      `gorm:"foreignKey:UserId;references:Id"`
	Orders   Orders     `gorm:"foreignKey:OrderId;references:Id"`
	IsActive bool       `gorm:"column:IsActive;"`
}

func (Sheet) TableName() string {
	return SheetTable
}
