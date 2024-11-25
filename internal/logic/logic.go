package logic

import (
	"context"
	"github/free-order-be/internal/client"
	"github/free-order-be/internal/dao"
	"github/free-order-be/models"
)

type Logic interface {
	//User
	GetUser(context.Context, *models.GetUserReq) (*models.User, error)
	CreateUser(context.Context, *models.User) (*models.User, error)
	SignIn(context.Context, *models.User) (*models.SignInResp, error)

	//Order
	CreateOrder(context.Context, *models.Order) (*models.Order, error)
	GetSheetOrders(context.Context, *models.GetSheetOrdersReq) (models.Orders, error)
	GetUserOrders(context.Context, *models.GetUserOrdersReq) (models.Orders, error)

	//Sheet
	CreateSheet(context.Context, *models.Sheet) (*models.Sheet, error)
	GetSheet(context.Context, *models.GetSheetReq) (*models.Sheet, error)

	//Shopee
	GetShopeeMenu(context.Context) (*client.GetDeliveryIDResp, error)
}
type LogicImpl struct {
	SecretKey []byte
	Client    *dao.DAO
	Shopee    *client.Shopee
}

func NewLogicImpl(dao *dao.DAO) *LogicImpl {
	return &LogicImpl{
		Client:    dao,
		SecretKey: []byte("Idasdasdasdnasdnjknxzm1323"),
	}
}
