package dao

import (
	"fmt"
	d "github/lambda-microservice/internal/domain"

	"gorm.io/gorm"
)

type IUserDAO interface {
	Create(*d.User) error
	Find(*d.User) (*d.User, error)
}

type UserImpl struct {
	client *gorm.DB
}

func (dao *UserImpl) Create(user *d.User) error {
	tx := dao.client.Create(user)
	if tx.Error != nil {
		return fmt.Errorf("internal error: %s", tx.Error)
	}
	return nil
}

func (dao *UserImpl) Find(req *d.User) (*d.User, error) {
	var result *d.User
	tx := dao.client.
		Table(d.UserTable).
		Where("UserName = ?", req.UserName).
		Preload("UserInfo").
		Find(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("internal error: %s", tx.Error)
	}
	return result.CheckNil(), nil
}
