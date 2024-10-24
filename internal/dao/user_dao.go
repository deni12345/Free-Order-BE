package dao

import (
	"context"
	d "github/free-order-be/internal/domain"

	"github.com/guregu/dynamo/v2"
)

const UserTable = "user"

type IUserDAO interface {
	Create(context.Context, *d.User) error
}

type UserImpl struct {
	client *dynamo.DB
	table  dynamo.Table
}

func (u *UserImpl) TableName() string {
	return u.table.Name()
}

func (u *UserImpl) Create(ctx context.Context, user *d.User) error {
	putRequest := u.table.Put(user)
	return putRequest.Run(ctx)
}

// func (u *UserImpl) Find(req *d.User) (*d.User, error) {
// 	var result *d.User
// 	tx := dao.client.
// 		Table(d.UserTable).
// 		Where("UserName = ?", req.UserName).
// 		Preload("UserInfo").
// 		Find(&result)
// 	if tx.Error != nil {
// 		return nil, fmt.Errorf("internal error: %s", tx.Error)
// 	}
// 	return result.CheckNil(), nil
// }
