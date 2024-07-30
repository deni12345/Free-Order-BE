package logic

import (
	"context"
	"github/lambda-microservice/internal/dao"
	"github/lambda-microservice/models"
)

type Logic interface {
	SignUp(*models.User) (*models.User, error)
	SignIn(context.Context, *models.User) (*models.SignInResp, error)
}
type LogicImpl struct {
	SecretKey []byte
	Client    *dao.DAO
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
		Client:    dao,
		SecretKey: []byte("Idasdasdasdnasdnjknxzm1323"),
	}
}
