package domain

// import (
// 	"time"
// )

// const (
// 	SheetTable = "Sheet"
// )

// type Sheets []*Sheet

// type Sheet struct {
// 	ID       *uint     `gorm:"column:Id;"`
// 	Name     string    `gorm:"column:Name;"`
// 	UserId   *uint     `gorm:"column:UserId;"`
// 	User     *User     `gorm:"foreignKey:UserId;references:Id"`
// 	Orders   Orders    `gorm:"foreignKey:SheetId;references:Id"`
// 	IsActive bool      `gorm:"column:IsActive;"`
// 	CreateAt time.Time `gorm:"column:CreateAt;"`
// }

// func (Sheet) TableName() string {
// 	return SheetTable
// }
