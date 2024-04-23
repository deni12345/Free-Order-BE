package domain

type OrderSheet struct {
	ID        *uint  `gorm:"column:Id;"`
	SheetName string `gorm:"column:SheetName;"`
	CreateAt  string `gorm:"column:CreateAt;"`
	UserId    *uint  `gorm:"column:UserId;"`
	User      *User  `gorm:"foreignKet:UserId;references:Id"`
	OrderId   *uint  `gorm:"column:OrderId;"`
	Order     *Order `gorm:"foreignKet:OrderId;references:Id"`
}
