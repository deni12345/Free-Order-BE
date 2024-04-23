package logic

import (
	"github/lambda-microservice/internal/dao"
	"github/lambda-microservice/internal/domain"
	"github/lambda-microservice/model"
)

type Logic interface {
	CreateUser(*model.User) (*model.User, error)
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

func (l *LogicImpl) CreateUser(req *model.User) (*model.User, error) {
	dUser := &domain.User{
		UserName: req.UserName,
		Password: req.Password,
	}
	err := l.Client.UserDAO.Create(dUser)
	if err != nil {
		return nil, err
	}
	return &model.User{
		UserName: dUser.UserName,
		Password: dUser.Password,
	}, nil
}
