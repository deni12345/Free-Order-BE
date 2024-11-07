package logic

import (
	"context"
	"fmt"
	"github/free-order-be/models"
)

func (l *LogicImpl) GetOrders(ctx context.Context, req *models.GetOrdersReq) (models.Orders, error) {
	if req == nil {
		return nil, fmt.Errorf("[Logic] Invalid get order request")
	}
	orders, err := l.Client.OrderDAO.FindsBySheet(ctx, req.GetSheetID())
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, fmt.Errorf("[Logic] there is no order founded")
	}
	return orders.GetModelOrders(), nil
}
