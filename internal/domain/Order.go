package domain

const (
	OrderTable = "User"
)

type Order struct {
	ID       *uint  `gorm:"column:Id;"`
	FoodName string `gorm:"column:FoodName;"`
	Amount   uint   `gorm:"column:Amount;"`
	CreateAt string `gorm:"column:Amount;"`
}

func (Order) TableName() string {
	return OrderTable
}
