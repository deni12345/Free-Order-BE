package logic

import (
	"context"
	"fmt"
	d "github/free-order-be/internal/domain"
	"github/free-order-be/models"
)

func (l *LogicImpl) CreateOrder(ctx context.Context, req *models.Order) (*models.Order, error) {
	ctxOrder := d.BuildDomainOrder(req)
	if ctxOrder == nil {
		return nil, fmt.Errorf("[Logic] cannot parse model order")
	}
	orders, err := l.Client.OrderDAO.FindsBySheet(ctx, ctxOrder.GetPK())
	if err != nil {
		return nil, err
	}
	if len(orders) > 0 {
		return nil, fmt.Errorf("order %v already exist", ctxOrder.GetName())
	}

	err = l.Client.OrderDAO.Create(ctx, ctxOrder)
	if err != nil {
		return nil, err
	}
	return ctxOrder.GetModelOrder(), nil
}
