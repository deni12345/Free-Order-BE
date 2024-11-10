package logic

import (
	"context"
	"fmt"
	"github/free-order-be/models"
)

func (l *LogicImpl) GetSheetOrders(ctx context.Context, req *models.GetSheetOrdersReq) (models.Orders, error) {
	if req == nil {
		return nil, fmt.Errorf("[Logic] Invalid get order request")
	}
	orders, err := l.Client.OrderDAO.FindAllBySheet(ctx, req.GetSheetID())
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, fmt.Errorf("[Logic] there is no order founded")
	}
	return orders.GetModelOrders(), nil
}
