package logic

import (
	"context"
	"fmt"
	d "github/free-order-be/internal/domain"
	"github/free-order-be/models"
)

func (l *LogicImpl) CreateOrder(ctx context.Context, req *models.Order) (*models.Order, error) {
	ctxOrder := d.BuildDomainOrder(req)
	if ctxOrder.IsValid() == nil {
		return nil, fmt.Errorf("[Logic] cannot parse model order")
	}
	order, err := l.Client.OrderDAO.FindByID(ctx, ctxOrder)
	if err != nil {
		return nil, err
	}
	if order.IsNil() {
		return nil, fmt.Errorf("order id %v already exist", ctxOrder.GetName())
	}

	err = l.Client.OrderDAO.Create(ctx, ctxOrder)
	if err != nil {
		return nil, err
	}
	return ctxOrder.GetModelOrder(), nil
}
