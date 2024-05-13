package logic

import (
	"github/lambda-microservice/internal/dao"
	"github/lambda-microservice/models"
)

type Logic interface {
	CreateUser(*models.User) (*models.User, error)
}
type LogicImpl struct {
	Client *dao.DAO
}

func NewLogicImpl() *LogicImpl {
	dao := dao.NewDAO(dao.Config{
		Port:     "3306",
		Host:     "localhost",
		User:     "root",
		Password: "password",
		DBName:   "fodb",
	})
	return &LogicImpl{
		Client: dao,
	}
}
