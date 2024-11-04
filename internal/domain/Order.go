package domain

import "time"

type Orders []*Order

type Order struct {
	PK       string    `dynamo:"PK,hash"`
	SK       string    `dynamo:"PK,hash"`
	Name     *uint     `gorm:"column:SheetId;"`
	UserId   *uint     `gorm:"foreignKey:UserId;references:Id"`
	FoodName string    `gorm:"column:FoodName;"`
	Amount   uint      `gorm:"column:Amount;"`
	CreateAt time.Time `gorm:"column:CreatAt;"`
}

func (Order) TableName() string {
	return OrderTable
}

// func (s *SheetImpl) CreateInfo(ctx context.Context, sheet *d.Sheet) error {
// 	newID, err := s.dao.NextID(ctx, SHEET_TABLE)
// 	if err != nil {
// 		return err
// 	}

// 	sheet.PK = fmt.Sprintf("SHEET#%v", *newID)
// 	sheet.SK = "INFO#METADATA"
// 	return s.table.Put(sheet).Run(ctx)
// }
