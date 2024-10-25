package logic

import (
	"context"
	"github/free-order-be/internal/dao"
	"github/free-order-be/models"
)

type Logic interface {
	GetUser(context.Context, *models.GetUserReq) (*models.User, error)
	SignUp(context.Context, *models.User) (*models.User, error)
	// SignIn(context.Context, *models.User) (*models.SignInResp, error)

	// CreateSheet(context.Context, *models.Sheet) (*models.Sheet, error)
	// UpdateSheet(context.Context, *models.Sheet) error
	//RemoveSheet(context.Context, *models.Sheet) error
}
type LogicImpl struct {
	SecretKey []byte
	Client    *dao.DAO
}

func NewLogicImpl(dao *dao.DAO) *LogicImpl {
	return &LogicImpl{
		Client:    dao,
		SecretKey: []byte("Idasdasdasdnasdnjknxzm1323"),
	}
}
