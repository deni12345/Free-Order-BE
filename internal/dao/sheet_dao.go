package dao

// import (
// 	"fmt"
// 	d "github/free-order-be/internal/domain"

// 	"github.com/guregu/dynamo/v2"
// )

// type ISheetDAO interface {
// 	Create(*d.Sheet) error
// 	//Find(*d.Sheet) (d.Sheets, error)
// }

// type SheetImpl struct {
// 	client *dynamo.DB
// }

// func (dao *SheetImpl) Create(sheet *d.Sheet) error {
// 	tx := dao.client.Create(sheet)
// 	if tx.Error != nil {
// 		return fmt.Errorf("internal error: %s", tx.Error)
// 	}
// 	return nil
// }

// func (dao *SheetImpl) Find(req *d.Sheet) (*d.User, error) {
// 	var result *d.User
// 	tx := dao.client.
// 		Table(d.UserTable).
// 		Where("Name = ?", req.Name).
// 		Preload("User", "Orders").
// 		Find(&result)
// 	if tx.Error != nil {
// 		return nil, fmt.Errorf("internal error: %s", tx.Error)
// 	}
// 	return result.CheckNil(), nil
// }
