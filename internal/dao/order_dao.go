package dao

import (
	"context"
	"fmt"
	d "github/free-order-be/internal/domain"

	"github.com/guregu/dynamo/v2"
)

type IOrderDAO interface {
	Create(context.Context, *d.Order) error
	FindByID(context.Context, *d.Order) (*d.Order, error)
	FindAllBySheet(context.Context, string) (d.Orders, error)
	FindAllByUser(context.Context, string) (d.Orders, error)
}

type OrderImpl struct {
	dao   *DAO
	table dynamo.Table
}

func NewOrderDAO(db *dynamo.DB) *OrderImpl {
	return &OrderImpl{
		dao:   NewDAORef(db),
		table: db.Table(SHEET_TABLE),
	}
}

func (o *OrderImpl) Create(ctx context.Context, order *d.Order) error {
	newID, err := o.dao.NextID(ctx, ORDER_TABLE)
	if err != nil {
		return err
	}
	if newID == nil {
		return fmt.Errorf("failed to get next id")
	}

	order.SK = o.createOrderSK(newID)
	return o.table.Put(order).Run(ctx)
}

func (o *OrderImpl) FindByID(ctx context.Context, order *d.Order) (*d.Order, error) {
	var result = &d.Order{}
	err := o.table.Get("PK", order.GetPK()).Range("SK", dynamo.Equal, order.GetSK()).One(ctx, result)
	if err != nil {
		if err == dynamo.ErrNotFound {
			return result, nil
		}
		return nil, err
	}
	return result, nil
}

func (o *OrderImpl) FindAllBySheet(ctx context.Context, sheetID string) (d.Orders, error) {
	var orders d.Orders
	err := o.table.Get("PK", sheetID).Range("SK", dynamo.BeginsWith, "ORDER#").All(ctx, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *OrderImpl) FindAllByUser(ctx context.Context, userID string) (d.Orders, error) {
	var orders d.Orders
	err := o.table.Scan().Index("UserOrderIndex").Filter("UserID=?", userID).All(ctx, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *OrderImpl) createOrderSK(id *uint) string {
	return fmt.Sprintf("ORDER#%v", *id)
}
