package logic

import (
	"context"
	"github/free-order-be/internal/dao"
	"github/free-order-be/models"
)

type Logic interface {
	//User
	GetUser(context.Context, *models.GetUserReq) (*models.User, error)
	SignUp(context.Context, *models.User) (*models.User, error)
	// SignIn(context.Context, *models.User) (*models.SignInResp, error)

	//Order
	CreateOrder(context.Context, *models.Order) (*models.Order, error)
	GetOrders(context.Context, *models.GetOrdersReq) (models.Orders, error)

	//Sheet
	CreateSheet(context.Context, *models.Sheet) (*models.Sheet, error)
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
