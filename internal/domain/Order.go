package domain

import "time"

const (
	OrderTable = "Order"
)

type Orders []*Order

type Order struct {
	ID       *uint     `gorm:"column:Id;"`
	SheetId  *uint     `gorm:"column:SheetId;"`
	UserId   *uint     `gorm:"foreignKey:UserId;references:Id"`
	FoodName string    `gorm:"column:FoodName;"`
	Amount   uint      `gorm:"column:Amount;"`
	CreateAt time.Time `gorm:"column:CreatAt;"`
}

func (Order) TableName() string {
	return OrderTable
}
