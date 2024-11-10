package logic

import (
	"context"
	"fmt"
	"github/free-order-be/models"
)

func (l *LogicImpl) GetUserOrders(ctx context.Context, req *models.GetUserOrdersReq) (models.Orders, error) {
	if req == nil {
		return nil, fmt.Errorf("[Logic] Invalid get user orders request")
	}
	orders, err := l.Client.OrderDAO.FindAllByUser(ctx, req.GetUserID())
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, fmt.Errorf("[Logic] there is no order founded")
	}
	return orders.GetModelOrders(), nil
}
