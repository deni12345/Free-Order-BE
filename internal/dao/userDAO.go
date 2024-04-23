package dao

import (
	"fmt"
	"github/lambda-microservice/internal/domain"

	"gorm.io/gorm"
)

type IUserDAO interface {
	Create(*domain.User) error
}

type UserImpl struct {
	client *gorm.DB
}

func (dao *UserImpl) Create(user *domain.User) error {
	tx := dao.client.Create(user)
	if tx.Error != nil {
		return fmt.Errorf("internal error: %s", tx.Error)
	}
	return nil
}
