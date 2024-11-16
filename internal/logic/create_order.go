package logic

import (
	"context"
	"fmt"
	"github/free-order-be/internal/domain"
	d "github/free-order-be/internal/domain"
	"github/free-order-be/models"
)

func (l *LogicImpl) CreateOrder(ctx context.Context, req *models.Order) (*models.Order, error) {
	ctxOrder := d.BuildDomainOrder(req)
	if !ctxOrder.IsValid() {
		return nil, fmt.Errorf("[Logic] cannot parse model order")
	}

	if order, err := l.Client.OrderDAO.FindByID(ctx, ctxOrder); err != nil {
		return nil, err
	} else if !order.IsNil() {
		return nil, fmt.Errorf("order id %v already exist", ctxOrder.GetName())
	}

	if err := l.Client.OrderDAO.Create(ctx, ctxOrder); err != nil {
		return nil, err
	}
	if err := l.Client.OrderDAO.CreateRealtime(ctx, domain.NewFirestoreOrder(ctxOrder)); err != nil {
		return nil, err
	}

	return ctxOrder.GetModelOrder(), nil
}
